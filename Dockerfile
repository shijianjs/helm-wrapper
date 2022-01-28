FROM golang:1.17

ENV GIN_MODE=release

COPY config.yaml  /app/config.yaml
COPY helm-wrapper /app/helm-wrapper
WORKDIR /app
USER root
CMD helm-wrapper
