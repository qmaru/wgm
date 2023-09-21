package logs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"wgm/common/utils"
	"wgm/common/configs"

	"github.com/sirupsen/logrus"
)

type myFormat struct{}

func (f *myFormat) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

// Logger 定义日志格式
func Logger(logName string) (*logrus.Logger, error) {
	logger := logrus.New()
	// 输出到文件
	if configs.Debug {
		logger.Out = os.Stdout
		logger.SetLevel(logrus.DebugLevel)
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
		accessFile, err := os.OpenFile(accessPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		logger.Out = accessFile
		logger.SetLevel(logrus.InfoLevel)
	}
	myformat := new(myFormat)
	logger.SetFormatter(myformat)
	return logger, nil
}
