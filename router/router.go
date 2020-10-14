// router 路由的初始化
package router

import (
    v1 "github.com/alonghub/GinBase/api/v1"
    "github.com/alonghub/GinBase/utils/log"
    "github.com/gin-gonic/gin"
    "time"
)

// StatCost 中间件，统计执行时间
func StatCost() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        // 可以设置一些公共参数
        c.Set("example", 123123)
        // 等待其他中间件执行
        c.Next()
        latency := time.Since(t)
        log.Logger.Debug(latency.String())
    }
}

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
    r := gin.Default()

    r.Use(StatCost())

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"result": "pong"})

    })

    r.HEAD("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"result": "pong"})
    })

    apiV1 := r.Group("/api/v1/sample")
    apiV1.Use()
    {
        apiV1.POST("/", v1.Sample)
    }

    return r
}
