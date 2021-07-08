# Introduction

Blazing fast, structured, leveled logging in Go.

# Features

- 高性能

# Usage


```go
package example

import (
	"errors"
	"testing"

	logger "github.com/huweihuang/logger/zap"
)

func TestZap(t *testing.T){
	c := logger.New()
	c.SetDivision("size")     // 设置归档方式，"time"时间归档 "size" 文件大小归档，文件大小等可以在配置文件配置
	c.SetTimeUnit(logger.Day) // 时间归档 可以设置切割单位
	c.SetEncoding("json")     // 输出格式 "json" 或者 "console"

	c.SetInfoFile("./logs/zap.log") // 设置info级别日志
	c.SetLogLevel("debug")
	c.InitLogger()

	logger.SugaredLogger.Info("info level test")
	logger.SugaredLogger.Error("error level test")
	logger.SugaredLogger.Warn("warn level test")
	logger.SugaredLogger.Debug("debug level test")

	logger.SugaredLogger.Infof("info level test: %s", "111")
	logger.SugaredLogger.Errorf("error level test: %s", "111")
	logger.SugaredLogger.Warnf("warn level test: %s", "111")
	logger.SugaredLogger.Debugf("debug level test: %s", "111")

	logger.Info("this is a log", logger.With("Trace", "12345677"))
	logger.Info("this is a log", logger.WithError(errors.New("this is a new error")))
}

func TestZapLogByToml(t *testing.T){
	c := logger.NewFromToml("../zap/config/config.toml")
	c.InitLogger()

	logger.SugaredLogger.Info("info level test")
	logger.SugaredLogger.Error("error level test")
	logger.SugaredLogger.Warn("warn level test")
	logger.SugaredLogger.Debug("debug level test")

	logger.SugaredLogger.Infof("info level test: %s", "111")
	logger.SugaredLogger.Errorf("error level test: %s", "111")
	logger.SugaredLogger.Warnf("warn level test: %s", "111")
	logger.SugaredLogger.Debugf("debug level test: %s", "111")

	logger.Info("this is a log", logger.With("Trace", "12345677"))
	logger.Info("this is a log", logger.WithError(errors.New("this is a new error")))
}
```
