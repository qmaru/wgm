package services

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func createFolder(folderPath string) string {
	_, err := os.Stat(folderPath)
	if err == nil {
		return folderPath
	}
	if os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
		return folderPath
	}
	return folderPath
}

func writeFile(f string, data []byte) {
	err := ioutil.WriteFile(f, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func MakeConfig(server, name, cfg string) {
	root, _ := os.Getwd()
	savePath := createFolder(filepath.Join(root, server))
	cfgPath := filepath.Join(savePath, name+".conf")
	writeFile(cfgPath, []byte(cfg))
}
