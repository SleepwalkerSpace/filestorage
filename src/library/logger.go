package library

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// NewLogger new logger
func NewLogger(path, name string, pretty bool) *logrus.Logger {
	logger := logrus.New()
	logger.AddHook(hook(path, name, pretty))
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
	return logger
}

func hook(path, name string, pretty bool) logrus.Hook {
	logrus.New()
	writer, err := rotatelogs.New(fmt.Sprintf("%s/%s", path, name)+".%Y%m%d%H", rotatelogs.WithRotationTime(time.Hour*24))
	if err != nil {
		logrus.Errorf("new logger %s/%s error: %v", path, name, err)
	}
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05", PrettyPrint: pretty})
}
