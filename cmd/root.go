package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var RootCmd = &cobra.Command{
	Use:     "GodzillaBatchAdd",
	Short:   "GodzillaBatchAdd",
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		GodzillaBatchAdd()
	},
}

var (
	id               string
	file             string
	password         string
	secretKey        string
	payload          string
	cryption         string
	encoding         string
	headers          string
	proxyType        string
	proxyHost        string
	proxyPort        int
	groupName        string
	databaseFilePath string
	createTime       string
	updateTime       string
	url              string
)

func init() {
	RootCmd.Flags().StringVarP(&file, "file", "f", "", "urls file path")
	RootCmd.Flags().StringVarP(&password, "password", "p", "pass", "password")
	RootCmd.Flags().StringVarP(&secretKey, "secretKey", "s", "key", "secretKey")
	RootCmd.Flags().StringVarP(&payload, "payload", "l", "JavaDynamicPayload", "")
	RootCmd.Flags().StringVarP(&cryption, "cryption", "c", "JAVA_AES_BASE64", "")
	RootCmd.Flags().StringVarP(&encoding, "encoding", "e", "UTF-8", "")
	RootCmd.Flags().StringVar(&headers, "headers", "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0", "")
	RootCmd.Flags().StringVar(&proxyType, "proxyType", "NO_PROXY", "")
	RootCmd.Flags().StringVar(&proxyHost, "proxyHost", "127.0.0.1", "")
	RootCmd.Flags().IntVar(&proxyPort, "proxyPort", 8888, "")
	RootCmd.Flags().StringVarP(&groupName, "groupName", "g", "/", "group name")
	RootCmd.Flags().StringVarP(&databaseFilePath, "databaseFilePath", "d", "data.db", "database file path")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GodzillaBatchAdd() {

	db, err := sql.Open("sqlite3", databaseFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if file == "" {
		fmt.Println("please input file path")
		os.Exit(1)
	}

	if exists := fileExists(file); exists == false {
		fmt.Println("file no found")
		os.Exit(1)
	}

	file, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url = scanner.Text()
		id = uuid.New().String()
		createTime = time.Now().Format("2006-01-02 15:04:05")
		updateTime = createTime

		_, err := db.Exec(`
			INSERT INTO shell (id, url, password, secretKey, payload, cryption, encoding, headers, reqLeft, reqRight, connTimeout, readTimeout, proxyType, proxyHost, proxyPort, remark, note, createTime, updateTime)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			id, url, password, secretKey, payload, cryption, encoding, headers, "", "", 3000, 6000, proxyType, proxyHost, proxyPort, "", "", createTime, updateTime)
		if err != nil {
			log.Fatal(err)
		}

		err = ceateGroup(db, id)
		if err != nil {
			log.Fatal(err)
		}

	}

}
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // 文件存在
	}
	if os.IsNotExist(err) {
		return false // 文件不存在
	}
	return false // 其他错误（比如权限问题）
}

func ceateGroup(db *sql.DB, id string) error {
	if groupName == "/" {
		return nil // 根目录就啥都不做
	}

	groupId := "/" + groupName

	// 1. 查询 group 是否存在
	var exists int
	err := db.QueryRow(`SELECT COUNT(*) FROM shellGroup WHERE groupId = ?`, groupId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("查询 shellGroup 出错: %w", err)
	}

	// 2. 如果不存在，先插入 group
	if exists == 0 {
		_, err := db.Exec(`INSERT INTO shellGroup (groupId) VALUES (?)`, groupId)
		if err != nil {
			return fmt.Errorf("插入 shellGroup 出错: %w", err)
		}
	}

	// 3. 插入 shellEnv
	_, err = db.Exec(`INSERT INTO shellEnv (shellId, key, value) VALUES (?, 'ENV_GROUP_ID', ?)`,
		id, groupId)
	if err != nil {
		return fmt.Errorf("插入 shellEnv 出错: %w", err)
	}

	return nil
}
