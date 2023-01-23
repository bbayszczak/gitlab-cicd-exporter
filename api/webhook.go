package api

import (
	"encoding/json"
	"net/http"

	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/bbayszczak/gitlab-cicd-exporter/logging"

	"github.com/labstack/echo/v4"
	gitlab "github.com/xanzy/go-gitlab"
)

func Webhook(c echo.Context) error {
	logger := logging.GetLoggerFromContext(c)
	cc := c.(*customcontext.CustomContext)
	pipelineEvent := gitlab.PipelineEvent{}
	err := cc.Context.Bind(&pipelineEvent)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if marshalled, err := json.Marshal(pipelineEvent); err != nil {
		logger.Warnw("cannot marshall gitlabEvent", "err", err)
	} else {
		logger.Debugw("gitlab event received", "event", string(marshalled))
	}
	return c.NoContent(http.StatusOK)
}
