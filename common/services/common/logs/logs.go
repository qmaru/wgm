package logs

import (
	"path/filepath"

	"wgm/common/configs"
	"wgm/common/utils"

	"github.com/gin-gonic/gin"
	"github.com/qmaru/qlog/ginlog"
)

func GinLogger(logName string) (gin.HandlerFunc, error) {
	var logFile string
	if configs.Debug {
		logFile = ""
	} else {
		logPath, err := utils.FileSuite.RootPath("logs")
		if err != nil {
			return nil, err
		}
		logpath, err := utils.FileSuite.Mkdir(logPath)
		if err != nil {
			return nil, err
		}
		accessPath := filepath.Join(logpath, logName)
		logFile = accessPath
	}

	return ginlog.Logger(logFile)
}
