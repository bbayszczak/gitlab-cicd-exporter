# gitlab-cicd-exporter

A Prometheus / OpenMetrics exporter for Gitlab CI/CD

This exporter is made to help to improve GitlabCI pipelines performance

It`s only using Gitlab webhooks and do not contact Gitlab API to avoid rate limiting issues

## table of content

- [gitlab-cicd-exporter](#gitlab-cicd-exporter)
  - [table of content](#table-of-content)
  - [configuration](#configuration)
  - [metrics list](#metrics-list)
  - [how to use](#how-to-use)
    - [run gitlab-cicd-exporter](#run-gitlab-cicd-exporter)
      - [Docker (recommended)](#docker-recommended)
      - [from source code](#from-source-code)
      - [Helm chart](#helm-chart)
    - [Gitlab setup](#gitlab-setup)

## configuration

| environment variable              | default value | possible values               | description                      |
|-----------------------------------|---------------|-------------------------------|----------------------------------|
| GITLAB_CICD_EXPORTER_GITLAB_TOKEN |               | any string                    | the token set to Gitlab webhooks |
| GITLAB_CICD_EXPORTER_LOG_LEVEL    | info          | (debug\|info\|warning\|error) | gitlab-cicd-exporter log level   |

## metrics list

| metric name                            | type      | description                    |
|----------------------------------------|-----------|--------------------------------|
| gitlab_cicd_pipelines_started_count    | counter   | the count of pipelines started |
| gitlab_cicd_pipelines_duration_seconds | histogram | pipelines duration in seconds  |
| gitlab_cicd_jobs_duration_seconds      | histogram | jobs duration in seconds       |

## how to use

There is a two part setup to use `gitlab-cicd-exporter`

- run `gitlab-cicd-exporter`

- setup Gitlab to send webhooks to your `gitlab-cicd-exporter` instance

### run gitlab-cicd-exporter

First of all you need to generate a token which will be used to secure incoming Gitlab Webhooks (`openssl rand -base64 24` can be useful :wink:)

Save this token for later use

All metrics are available under the `/metrics` endpoint

#### Docker (recommended)

`docker run -p 8080:8080 -e GITLAB_CICD_EXPORTER_GITLAB_TOKEN=<VALUE PREVIOUSLY GENERATED> ghcr.io/bbayszczak/gitlab-cicd-exporter:latest`

This will start `gitlab-cicd-exporter` and make it listen on port `8080`

You can visualize metrics with `curl -s http://localhost:8080/metrics`

You can use `ngrok` if you want to expose `gitlab-cicd-exporter` to Gitlab (`ngrok http 8080`)

#### from source code

1. clone the repository

2. run a `make build` to generate a `gitlab-cicd-exporter` binary

3. export the configuration environment variables you want (at least the `GITLAB_CICD_EXPORTER_GITLAB_TOKEN` variable with the previously generated token)

4. run the binary `gitlab-cicd-exporter`: it will start the server on port `8080`

You can visualize metrics with `curl -s http://localhost:8080/metrics`

You can use `ngrok` if you want to expose `gitlab-cicd-exporter` to Gitlab (`ngrok http 8080`)

#### Helm chart

TO BE DONE

### Gitlab setup

A webhook has to be setup on Gitlab side

It can be done at a group or a project level

- project level: only project pipelines metrics will be exported

- group level: all projects under the group will have their metrics exported

1. Go to `Settings`>`Webhooks` for your group/project

2. Create a new Webhook

    **URL**: URL targeting your `gitlab-cicd-exporter` instance and add `/webhook` and the end

    eg. `https://146f-2a01-e0a-3da-44c1-c4f1-ac40-5339-b5f3.eu.ngrok.io/webhook`

    **Secret token**: the token generated previously

    **Trigger**: select the `Pipeline events` checkbox

    Then click on `Add webhook`

Next time pipeline is trigerred, `gitlab-cicd-exporter` will receive a webhook
