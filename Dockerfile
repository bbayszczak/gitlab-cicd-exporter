FROM golang:1.19-alpine as builder

WORKDIR /build

RUN apk add --no-cache make

RUN adduser -u 10001 --disabled-password scratchuser

COPY . /build

RUN CGO_ENABLED=0 GOOS=linux make build-release

FROM scratch

EXPOSE 8080/tcp

WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /build/gitlab-cicd-exporter /app

USER scratchuser

ENTRYPOINT ["/app/gitlab-cicd-exporter"]