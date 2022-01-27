FROM golang:1.17-alpine

ENV GIN_MODE=release

COPY config.yaml  /app/config.yaml
COPY helm-wrapper /app/helm-wrapper

CMD [ "/app/helm-wrapper" ]
