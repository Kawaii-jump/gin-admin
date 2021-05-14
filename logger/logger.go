package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Fileds warps logrus.Fileds ,type is map[string]interface{}
type Fileds logrus.Fields

// WriterHook is a log hook
type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

// Fire ...
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	// _ -> int型数据,返回写入数据的字节数
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels ...
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	logger.Level = logrus.InfoLevel
	logger.SetFormatter(&logrus.TextFormatter{})
	appName := os.Getenv("TCE_SERVICE_NAME")
	if appName == "" {
		appName = "app"
	}
	logPath := "log"
	if _, err := os.Stat(logPath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(logPath, os.ModePerm)
		}
	}
	os.Mkdir("log", os.ModePerm)
	if os.Getenv("TCE_HOST_ENV") == "online" {
		logPath = "/opt/tiger/toutiao/log/app"
		_, err := os.Stat(logPath)
		if err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(logPath, os.ModePerm)
			}
		}
	}
	errorLogName := fmt.Sprintf("%s/%s_%s.log", logPath, appName, "error")
	infoLogName := fmt.Sprintf("%s/%s_%s.log", logPath, appName, "info")
	warningLogName := fmt.Sprintf("%s/%s_%s.log", logPath, appName, "warning")

	var err error
	var infoOpenFile *os.File
	var errorOpenFile *os.File
	var warningOpenFile *os.File

	errorOpenFile, err = os.OpenFile(errorLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	checkError(err)
	infoOpenFile, err = os.OpenFile(infoLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	checkError(err)
	warningOpenFile, err = os.OpenFile(warningLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	checkError(err)

	logger.SetOutput(ioutil.Discard) // Send all logs to nowhere by default
	logger.AddHook(&WriterHook{      // Send logs with level higher than warning to stderr
		Writer: warningOpenFile,
		LogLevels: []logrus.Level{
			logrus.WarnLevel,
		},
	})
	logger.AddHook(&WriterHook{ // Send logs with level higher than warning to stderr
		Writer: errorOpenFile,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
	})
	logger.AddHook(&WriterHook{ // Send info and debug logrus. to stdout
		Writer: infoOpenFile,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})
}
