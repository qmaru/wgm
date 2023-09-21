package configs

import (
	"path/filepath"
	"runtime"

	"wgm/common/utils"

	"github.com/adrg/xdg"
)

const databaseName = "wgm.db"

// DatabaseConfig 数据库连接配置
func DatabaseConfig() (string, error) {
	if runtime.GOOS == "windows" {
		mainRoot, err := utils.FileSuite.RootPath("")
		if err != nil {
			return "", err
		}
		cfgPath := filepath.Join(mainRoot, databaseName)
		return cfgPath, nil
	}
	cfgRoot := filepath.Join(xdg.UserDirs.Documents, "wgm")
	cfgRootF, err := utils.FileSuite.Mkdir(cfgRoot)
	if err != nil {
		return "", err
	}
	cfgPath := filepath.Join(cfgRootF, databaseName)
	return cfgPath, nil
}
