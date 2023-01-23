package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/bbayszczak/gitlab-cicd-exporter/eventhandlers"
	"github.com/bbayszczak/gitlab-cicd-exporter/logging"
	"github.com/bbayszczak/gitlab-cicd-exporter/utils"

	"github.com/labstack/echo/v4"
	gitlab "github.com/xanzy/go-gitlab"
)

func Webhook(c echo.Context) error {
	logger := logging.GetLoggerFromContext(c)
	cc := c.(*customcontext.CustomContext)
	gitlabEventType := cc.Context.Request().Header.Get("X-Gitlab-Event")
	logger.Debugw("webhook event type", "gitlabEventType", gitlabEventType)
	rawBody, err := ioutil.ReadAll(cc.Context.Request().Body)
	if err != nil {
		logger.Errorw("cannot get raw body from request", "err", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	gitlabEvent, err := gitlab.ParseWebhook(gitlab.EventType(gitlabEventType), rawBody)
	if err != nil {
		logger.Errorw("cannot parse Gitlab event", "err", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// DEBUG, marshall to file
	if marshalled, err := json.Marshal(gitlabEvent); err != nil {
		logger.Warnw("cannot marshall gitlabEvent", "err", err)
	} else {
		logger.Debugw("gitlab event received", "event", string(marshalled))
		if err := utils.BytesToFile("events", "event", marshalled); err != nil {
			logger.Warnw("cannot save event to file", "err", err)
		}
	}

	switch event := gitlabEvent.(type) {
	case *gitlab.PipelineEvent:
		err = eventhandlers.PipelineEvent(event, logger, cc)
	default:
		logger.Warnw("unknown event received", "eventType", reflect.TypeOf(event))
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
