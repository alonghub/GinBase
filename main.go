// main 主函数
package main

import (
    "fmt"
    "github.com/alonghub/GinBase/router"
    "github.com/alonghub/GinBase/utils/cache"
    "github.com/alonghub/GinBase/utils/config"
    "github.com/alonghub/GinBase/utils/log"
    "net/http"
)

// main 函数入口
func main() {
    // 初始化配置文件
    config.Setup()
    // 初始化日志相关内容
    log.ZapLoggerInit("debug", "GinBase")
    // 初始化Redis连接池
    cache.RedisPoolInit()
    // 路由初始化
    routersInit := router.InitRouter()

    server := &http.Server{
        Handler: routersInit,
        Addr:    fmt.Sprintf("0.0.0.0:%d", config.ServerSetting.HttpPort),
    }
    err := server.ListenAndServe()
    if err != nil {
        log.Logger.Fatal(err.Error())
    }

}
