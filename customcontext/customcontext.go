package customcontext

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/configuration"
	"github.com/bbayszczak/gitlab-cicd-exporter/metrics"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CustomContext struct {
	echo.Context
	Log     *zap.SugaredLogger
	Config  *configuration.Configuration
	Metrics *metrics.Metrics
}
