package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Garfield247/gonf.git/config"
)

// 获取Log文件路径
func GetLogFileFullPath() string {
	prefixPath := config.LogCfg.LogSavePath
	suffixPath := fmt.Sprintf(
		"%s%s.%s",
		config.LogCfg.LogSaveName,
		time.Now().Format(config.LogCfg.TimeFormat),
		config.LogCfg.LogFileExt,
	)

	fullPath := filepath.Join(prefixPath, suffixPath)
	fmt.Println(fullPath)
	return fullPath
}

// 开启Log文件
func openLogFile(filePath string) *os.File {
	// 判断文件路径是否存在
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		dir, _ := os.Getwd()
		absPath := filepath.Join(dir, filePath)
		errr := os.MkdirAll(absPath, os.ModePerm)
		if errr != nil {
			panic(errr)
		}
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Field to OpenFile :%v", err)
	}
	return handle
}
