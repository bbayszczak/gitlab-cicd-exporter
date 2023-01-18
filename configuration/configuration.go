package configuration

import (
	"os"

	"go.uber.org/zap"
)

type Configuration struct {
	GitlabToken string `yaml:"gitlab_token"`
}

func GetConfiguration(sugaredLogger *zap.SugaredLogger) *Configuration {
	sugaredLogger.Info("reading configuration")
	gitlabToken := os.Getenv("GPM_GITLAB_TOKEN")
	if gitlabToken == "" {
		sugaredLogger.Warn("the GPM_GITLAB_TOKEN environment variable is empty")
	}
	return &Configuration{
		GitlabToken: gitlabToken,
	}
}
