package api

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/bbayszczak/gitlab-cicd-exporter/logging"
	"github.com/labstack/echo/v4"
)

func GitlabAuth(auth string, c echo.Context) (bool, error) {
	logger := logging.GetLoggerFromContext(c)
	cc := c.(*customcontext.CustomContext)
	logger.Debugw("gitlab token auth", "GitlabToken", cc.Config.GitlabToken, "EventToken", auth)
	return cc.Config.GitlabToken == auth, nil
}
