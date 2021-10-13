package apis

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"wgm/assets"
	"wgm/config"
	"wgm/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type myLoggerFormat struct{}

func (f *myLoggerFormat) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	logger := logrus.New()

	// 输出到文件
	if config.Debug {
		logger.Out = os.Stdout
		logger.SetLevel(logrus.DebugLevel)
	} else {
		currentPath := utils.FileSuite.LocalPath(config.Debug)
		logPath := filepath.Join(currentPath, "logs")
		logpath := utils.FileSuite.Create(logPath)
		accessPath := filepath.Join(logpath, "access.log")
		accessFile, _ := os.OpenFile(accessPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend)

		logger.Out = accessFile
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.SetFormatter(new(myLoggerFormat))

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
	}
}

// Run 执行服务
func Run() {
	listenAddr := "127.0.0.1:8373"

	log.Println("Listen: " + listenAddr)

	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	// 跨域
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTION"}
	router.Use(cors.New(config))

	router.Use(assets.StaticHand())
	router.Use(gin.Recovery())
	router.Use(Logger())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/server/:serverID", ServerInfo)
		v1.POST("/server", ServerAdd)
		v1.PUT("/server/:serverID", ServerUpdate)
		v1.DELETE("/server/:serverID", ServerDel)

		v1.GET("/node/:serverID/:userID", NodeInfo)
		v1.POST("/node/:serverID", NodeAdd)
		v1.PUT("/node/:serverID/:userID", NodeUpdate)
		v1.PUT("/node/:serverID/:userID/key", NodeUpdateKey)
		v1.DELETE("/node/:serverID/:userID", NodeDel)

		v1.POST("/rule", RuleAdd)
		v1.PUT("/rule/:ruleID", RuleUpdate)
		v1.DELETE("/rule/:ruleID", RuleDel)

		v1.POST("/rulemap/:ruleID/:userID", RuleMapAdd)
		v1.DELETE("/rulemap/:ruleID/:userID", RuleMapDel)

		v1.GET("/data/rulelist", RuleList)
		v1.GET("/data/serverlist", ServerList)
		v1.GET("/data/nodelist", NodeList)
		v1.GET("/data/servergroup", ServerGroup)
		v1.GET("/data/config", ShowConfig)
		v1.PUT("/data/userrule/:userID", UserRuleSets)
	}

	router.Run(listenAddr)
}
