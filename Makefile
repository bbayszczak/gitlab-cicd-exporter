NAME				= gitlab-cicd-exporter

SRCS				= main.go

GO_RUN				= go run
GO_BUILD			= go build -o $(NAME)

RM					= rm -rf

ENV_VARS			= GITLAB_CICD_EXPORTER_LOG_LEVEL=debug GITLAB_CICD_EXPORTER_GITLAB_TOKEN=test

.PHONY: all run clean

all: run

run:
	$(ENV_VARS) $(GO_RUN) $(SRCS)

clean:
	$(RM) $(NAME)
