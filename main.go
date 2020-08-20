package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/majordomo/routers"
	"github.com/guoruibiao/majordomo/scripts"
	"log"
)

func init() {
	// do
}

func main() {
    router := gin.Default()
    // 静态文件路由设置
    router.Static("/statics/", "./statics")
    router.Static("/templates/", "./templates")

    // 动态注册服务路由
    routers.Init(router)

    // 注册监听事件
    scripts.StartClipboardListening()

    // 服务启动
    err := router.Run(":8080")
    if err != nil {
	log.Fatal(err)
    }
}
