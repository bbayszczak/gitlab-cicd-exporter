package eventhandlers

import (
	"github.com/bbayszczak/gitlab-cicd-exporter/customcontext"
	"github.com/bbayszczak/gitlab-cicd-exporter/metrics"
	gitlab "github.com/xanzy/go-gitlab"
	"go.uber.org/zap"
)

func PipelineEvent(event *gitlab.PipelineEvent, logger *zap.SugaredLogger, cc *customcontext.CustomContext) error {
	// if pipeline started
	if event.ObjectAttributes.Status == "running" {
		logger.Debug("pipeline started identified")
		metrics.CounterIncrease(cc.Metrics.PipelinesStartedCount, map[string]string{
			"project": event.Project.PathWithNamespace,
			"source":  event.ObjectAttributes.Source,
			"ref":     event.ObjectAttributes.Ref,
		})
	}

	// if pipeline ended
	if event.ObjectAttributes.Status != "running" && event.ObjectAttributes.Status != "pending" {
		logger.Debug("pipeline ended identified")
		metrics.HistogramObserve(cc.Metrics.PipelinesDuration, float64(event.ObjectAttributes.Duration), map[string]string{
			"project": event.Project.PathWithNamespace,
			"source":  event.ObjectAttributes.Source,
			"ref":     event.ObjectAttributes.Ref,
			"status":  event.ObjectAttributes.Status,
		})
	}
	return nil
}
