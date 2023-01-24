# gitlab-cicd-exporter

A Prometheus / OpenMetrics exporter for Gitlab CI/CD

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
