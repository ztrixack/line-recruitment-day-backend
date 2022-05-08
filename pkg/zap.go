package pkg

import (
	"regexp"

	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func InitLog() {
	logger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		Exit(err)
	}
	defer logger.Sync()
	Sugar = logger.Sugar()
}

func Info(typ string, format string, args ...interface{}) {
	Sugar.Infof("["+typ+"] "+format, args...)
}

func Error(err error, format string, args ...interface{}) {
	Sugar.Debugf("[XX] "+format, args...)
	re := regexp.MustCompile(`\r?\n`)
	Sugar.Debugf("[__] " + re.ReplaceAllString(err.Error(), " "))
	// Sugar.Error(err) // for debug
}

func Warp(txt string, num int) string {
	return txt[len(txt)-6:]
}
