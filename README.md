# majordomo

监听系统剪切板，记录到 elasticsearch 中，做后置内容查找分析使用。

## 使用须知

1. 安装 elasticsearch
```shell_script
https://www.elastic.co/cn/downloads/elasticsearch
```

使用 elasticsearch 须安装 Java8 及以上版本，配置好 JAVA_HOME 等环境变量。

2. 使用 majordomo

下载本 repo，直接运行 main.go 文件,会自动在后台监听系统剪切板变化。
```go
go run main.go
```

3. 查询历史内容

```shell_script
http://localhost:8080/ 首页
http://localhost:8080/templates/clipboard.html 功能页
```

Enjoy.