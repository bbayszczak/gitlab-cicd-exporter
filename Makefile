NAME				= gitlab-cicd-exporter

SRCS				= main.go

GO_RUN				= go run
GO_BUILD			= go build -o $(NAME)
GO_BUILD_ARCH		= GOARCH=amd64 CGO_ENABLED=0 GOOS=linux
GO_BUILD_LDFLAGS	= -ldflags="-w -s"

RM					= rm -rf

ENV_VARS			= GITLAB_CICD_EXPORTER_LOG_LEVEL=debug GITLAB_CICD_EXPORTER_GITLAB_TOKEN=test

.PHONY: all run clean

all: run

run:
	$(ENV_VARS) $(GO_RUN) $(SRCS)

build: clean
	$(GO_BUILD) $(SRCS)

build-release: clean
	$(GO_BUILD_ARCH) $(GO_BUILD) $(GO_BUILD_LDFLAGS)

clean:
	$(RM) $(NAME)
