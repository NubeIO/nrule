package logger

import (
	"fmt"
	"github.com/NubeIO/nrule/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"strings"
)

var Logger Log

type Log struct {
	*logrus.Logger
}

func Init() {
	Logger.Logger = New()
}

func New() *logrus.Logger {
	logLevel := viper.GetString("server.log.level")
	logrusLevel := logrus.InfoLevel
	switch logLevel {
	case "PANIC":
		logrusLevel = logrus.PanicLevel
	case "FATAL":
		logrusLevel = logrus.FatalLevel
	case "ERROR":
		logrusLevel = logrus.ErrorLevel
	case "WARN":
		logrusLevel = logrus.WarnLevel
	case "DEBUG":
		logrusLevel = logrus.DebugLevel
	case "TRACE":
		logrusLevel = logrus.TraceLevel

	}
	logger := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
			DisableColors: false,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				return fmt.Sprintf("%s()>", funcName), fmt.Sprintf(" %s:%d", filename, f.Line)
			},
		},
		Hooks: make(logrus.LevelHooks),
		Level: logrusLevel,
	}
	logger.SetReportCaller(true)
	if viper.GetBool("server.log.store") {
		file := fmt.Sprintf("%s/rubix-edge-wires.log", config.Config.GetAbsDataDir())
		fileHook, err := NewLogrusFileHook(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.Hooks.Add(fileHook)
		}
	}
	return logger
}
