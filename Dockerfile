FROM alpine:3.3

ENV RESULT_SERVICE_DIR /go/src/github.com/gaia-docker/tugbot-result-service-es

WORKDIR $RESULT_SERVICE_DIR

ADD .dist/tugbot-result-service-es $RESULT_SERVICE_DIR/tugbot-result-service-es

CMD ["./tugbot-result-service-es"]