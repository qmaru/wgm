package logs

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger(logName string) (gin.HandlerFunc, error) {
	logger, err := Logger(logName)
	if err != nil {
		return nil, err
	}
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("- %s %d %s %s %s",
			reqMethod,
			statusCode,
			clientIP,
			reqURI,
			latencyTime,
		)
	}, nil
}
