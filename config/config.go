package config

import (
	"path/filepath"

	"wgm/utils"
)

func SetDBPath(name string) string {
	currentPath := utils.FileSuite.LocalPath(Debug)
	dbfile := filepath.Join(currentPath, name)
	return dbfile
}
