package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("Cannot init logger: %s", err))
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	sugaredLogger := logger.Sugar()

	sugaredLogger.Info("starting gitlab-cicd-exporter")
}
