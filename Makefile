NAME				= gitlab-cicd-exporter

SRCS				= main.go

GO_RUN				= go run
GO_BUILD			= go build -o $(NAME)

RM					= rm -rf

.PHONY: all run clean

all: run

run:
	$(GO_RUN) $(SRCS)

clean:
	$(RM) $(NAME)
