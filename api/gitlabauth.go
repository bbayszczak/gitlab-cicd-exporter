package api

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/labstack/echo/v4"
)

func GitlabAuth(auth string, c echo.Context) (bool, error) {
	cc := c.(*customcontext.CustomContext)
	return cc.Config.GitlabToken == auth, nil
}
