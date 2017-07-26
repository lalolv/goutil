package goutil

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 创建日志文件
// 保存到当前目录下的 log 子目录中
func Log(info interface{}, prefix string) {
	// 当前工作路径
	pwd, _ := os.Getwd()
	filename := time.Now().Format("2006-01-02") + ".log"
	filepath := pwd + "/log/" + filename
	logfile, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}

	defer logfile.Close()
	// log.SetOutput(logfile)
	logger := log.New(logfile, "["+prefix+"] ", log.Ltime)
	logger.Print(info)
}
