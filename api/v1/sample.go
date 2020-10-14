// v1 第一版本的相关功能函数
package v1

import (
    "github.com/alonghub/GinBase/common"
    "github.com/alonghub/GinBase/utils/log"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "net/http"
)

// Sample 示例函数
func Sample(c *gin.Context) {
    var reqInfo common.ReqInfo
    if err := c.BindJSON(&reqInfo); err != nil {
        log.Logger.Error("param is wrong.")
        c.JSON(http.StatusBadRequest, gin.H{"msg": "error params"})
        return
    }
    logReqInfo := zap.Any("Sample", reqInfo)
    log.Logger.Info("Sample", logReqInfo)
    c.JSON(http.StatusOK, gin.H{"code": 10000, "data": reqInfo})
}
