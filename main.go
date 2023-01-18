package main

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/configuration"
	"github.com/bbayszczak/gitlab-cicd-exporter/logging"
	"go.uber.org/zap"
)

func main() {
	logger, loggerAtomLvl := logging.InitLogger()

	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	sugaredLogger := logger.Sugar()

	sugaredLogger.Info("starting gitlab-cicd-exporter")
	config := configuration.GetConfiguration(sugaredLogger)
	logging.SetLogLevel(sugaredLogger, loggerAtomLvl, config.LogLevel)
}
