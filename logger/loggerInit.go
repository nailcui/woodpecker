package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger
var WebLog *logrus.Logger

func init() {
	initAppLogger()
	initWebLogger()
}

func initAppLogger() {
	logFileName := "woodpecker.log"
	Logger = initLogger(logFileName)
}

func initWebLogger() {
	logFileName := "web.log"
	WebLog = initLogger(logFileName)
}

func initLogger(logFileName string) *logrus.Logger {

	logPath := "/logs/woodpecker"
	fileName := path.Join(logPath, logFileName)
	_, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("OpenFile err", err)
	}
	log := logrus.New()

	// 默认日志级别 info
	defaultLevel := logrus.InfoLevel
	level := "info"
	if level != "" {
		parseLevel, err := logrus.ParseLevel(level)
		if err != nil {
			fmt.Println("parseLevel err", err)
		} else {
			defaultLevel = parseLevel
		}
	}
	log.Level = defaultLevel

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writers := []io.Writer{
		logWriter,
		os.Stdout}
	if !true {
		writers = []io.Writer{logWriter}
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	//log.Formatter = &logrus.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//}
	return log
}
