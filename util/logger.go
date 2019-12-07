package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

//Logger - A logrus logger instance
var Logger *log.Logger

func init() {
	Logger = log.New()
	Logger.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(log.InfoLevel)
}
