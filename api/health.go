package api

import (
	"net/http"

	"github.com/bbayszczak/gitlab-cicd-exporter/logging"
	"github.com/labstack/echo/v4"
)

func Health(c echo.Context) error {
	logger := logging.GetLoggerFromContext(c)
	logger.Info("tototot")
	return c.String(http.StatusOK, "healthy")
}
