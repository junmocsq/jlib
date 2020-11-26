package jtools

import (
	"bufio"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// "/Users/junmo/go/src/jlib/logs"
func SetLogs(logPath string) {
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	path := logPath + "/jlogs.log"
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err != nil {
		logrus.Fatal("初始化hook失败")
	}

	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		},
		&logrus.TextFormatter{},
	))

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.Fatal("打开/dev/null 失败")
	}
	nullWriter := bufio.NewWriter(src)
	logrus.SetOutput(nullWriter)

}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
