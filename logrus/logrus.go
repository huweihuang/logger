package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

const (
	defaultLevel = "info"
)

func InitLogger(logFile, logLevel, format string, enableReportCaller, enableForceColors bool) *logrus.Logger {
	logger := logrus.New()

	// set log level
	if logLevel == "" {
		logLevel = defaultLevel
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic("Failed to parse log level")
	}
	logger.SetLevel(level)

	// set logfile
	if logFile == "" {
		logger.SetOutput(os.Stdout)
	} else {
		accessLog, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to create access.log")
		}
		logger.SetOutput(accessLog)
	}

	// set log format
	forceColors := false
	if enableForceColors {
		forceColors = true
	}
	switch format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   forceColors,
		})
	}

	// set report caller
	if enableReportCaller {
		logger.SetReportCaller(true)
	}

	Logger = logger
	return logger
}
