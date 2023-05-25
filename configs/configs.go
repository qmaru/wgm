package configs

import (
	"path/filepath"

	"wgm/utils"
)

const databaseName = "wgm.db"

// DatabaseConfig 数据库连接配置
func DatabaseConfig() (string, error) {
	mainRoot, err := utils.FileSuite.RootPath("")
	if err != nil {
		return "", err
	}
	cfgPath := filepath.Join(mainRoot, databaseName)
	return cfgPath, nil
}
