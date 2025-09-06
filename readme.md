

由于哥斯拉是将连接的配置信息放在data.db数据库中的，所以可以直接操作data.db进行批量添加shell

```
GodzillaBatchAdd

Usage:
  GodzillaBatchAdd [flags]

Flags:
  -c, --cryption string            (default "JAVA_AES_BASE64")
  -d, --databaseFilePath string   database file path (default "data.db")
  -e, --encoding string            (default "UTF-8")
  -f, --file string               urls file path
  -g, --groupName string          group name (default "/")
      --headers string             (default "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")
  -h, --help                      help for GodzillaBatchAdd
  -p, --password string           password (default "pass")
  -l, --payload string             (default "JavaDynamicPayload")
      --proxyHost string           (default "127.0.0.1")
      --proxyPort int              (default 8888)
      --proxyType string           (default "NO_PROXY")
  -s, --secretKey string          secretKey (default "key")
  -v, --version                   version for GodzillaBatchAdd
```



使用方法

```
 ./godzillaBatchAdd -f webshell.txt -d ~/tools/jar-tools/data.db  -p your_pass-s your_key -g newGroup
```

![](https://cdn.jsdelivr.net/gh/zIxyd/image@main/ctfshow/20250906124310433.png)