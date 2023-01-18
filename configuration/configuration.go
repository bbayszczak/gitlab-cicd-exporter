package configuration

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Configuration struct {
	GitlabToken string
	LogLevel    zapcore.Level
}

func GetConfiguration(sugaredLogger *zap.SugaredLogger) *Configuration {
	sugaredLogger.Info("reading configuration")
	gitlabToken := os.Getenv("GITLAB_CICD_EXPORTER_GITLAB_TOKEN")
	if gitlabToken == "" {
		sugaredLogger.Warn("the GITLAB_CICD_EXPORTER_GITLAB_TOKEN environment variable is empty")
	}
	var zapLvl zapcore.Level
	switch os.Getenv("GITLAB_CICD_EXPORTER_LOG_LEVEL") {
	case "debug":
		zapLvl = zapcore.DebugLevel
	case "info":
		zapLvl = zapcore.InfoLevel
	case "warning":
		zapLvl = zapcore.WarnLevel
	case "error":
		zapLvl = zapcore.ErrorLevel
	default:
		zapLvl = zapcore.InfoLevel
	}
	config := &Configuration{
		GitlabToken: gitlabToken,
		LogLevel:    zapLvl,
	}
	sugaredLogger.Debugw("configuration read", "configuration", config)
	return config
}
