package eventhandlers

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/bbayszczak/gitlab-cicd-exporter/metrics"
	gitlab "github.com/xanzy/go-gitlab"
	"go.uber.org/zap"
)

func PipelineEvent(event *gitlab.PipelineEvent, logger *zap.SugaredLogger, cc *customcontext.CustomContext) error {
	if event.ObjectAttributes.Status == "running" {
		metrics.CounterIncrease(cc.Metrics.PipelinesStartedCount, map[string]string{"project": event.Project.PathWithNamespace})
	}
	return nil
}
