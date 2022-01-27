FROM alpine:3

ENV GIN_MODE=release

COPY config.yaml  /app/config.yaml
COPY helm-wrapper /app/helm-wrapper

CMD [ "/app/helm-wrapper" ]
