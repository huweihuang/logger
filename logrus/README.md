# Introduction

Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.

# Features

- 方便易用

# Usage

```go
package example

import (
	"testing"

	"github.com/sirupsen/logrus"

	logger "github.com/huweihuang/logger/logrus"
)

func TestLogrus(t *testing.T) {
	// init logger
	logger.InitLogger("./logs/logrus.log", "debug", "text", false, false)

	// Printf
	logger.Logger.Debugf("test debugf, %s", "debugf")
	logger.Logger.Infof("test infof, %s", "infof")
	logger.Logger.Warnf("test warnf, %s", "warnf")
	logger.Logger.Errorf("test errorf, %s", "errorf")

	// WithField
	logger.Logger.WithField("field1", "debug").Debug("test field, debug")
	logger.Logger.WithField("field1", "info").Info("test field, info")
	logger.Logger.WithField("field1", "warn").Warn("test field, warn")
	logger.Logger.WithField("field1", "error").Error("test field, error")

	// WithFields
	logger.Logger.WithFields(logrus.Fields{
		"fields1": "fields1_value",
		"fields2": "fields2_value",
	}).Debug("test fields, debug")

	logger.Logger.WithFields(logrus.Fields{
		"fields1": "fields1_value",
		"fields2": "fields2_value",
	}).Info("test fields, info")

	logger.Logger.WithFields(logrus.Fields{
		"fields1": "fields1_value",
		"fields2": "fields2_value",
	}).Warn("test fields, warn")

	logger.Logger.WithFields(logrus.Fields{
		"fields1": "fields1_value",
		"fields2": "fields2_value",
	}).Error("test fields, error")
}
```
