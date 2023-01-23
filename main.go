package main

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/api"
	"github.com/bbayszczak/gitlab-cicd-exporter/configuration"
	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/bbayszczak/gitlab-cicd-exporter/logging"
	"github.com/bbayszczak/gitlab-cicd-exporter/metrics"
	"go.uber.org/zap"

	echoprom "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	logger, loggerAtomLvl := logging.InitLogger()

	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	sugaredLogger := logger.Sugar()

	sugaredLogger.Info("starting gitlab-cicd-exporter")
	config := configuration.GetConfiguration(sugaredLogger)
	logging.SetLogLevel(sugaredLogger, loggerAtomLvl, config.LogLevel)

	metrics := metrics.NewMetrics()
	e := echo.New()
	e.HideBanner = true
	p := echoprom.NewPrometheus("echo", nil, metrics.MetricList())
	p.Use(e)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customcontext.CustomContext{
				Context: c,
				Log:     logger.Sugar(),
				Config:  config,
				Metrics: metrics,
			}
			return next(cc)
		}
	})
	e.Use(middleware.RequestID())
	e.Use(logging.ZapLogger(logger))
	e.GET("/health", api.Health)
	e.POST("/webhook", api.Webhook, middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:X-Gitlab-Token",
		AuthScheme: "",
		Validator:  api.GitlabAuth,
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
