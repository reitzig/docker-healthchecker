FROM alpine:3.10.1 AS run

ARG checks
ENV CHECKS=${checks}

RUN apk update && apk add --no-cache \
    postgresql-client=11.4-r0 \
    curl=7.65.1-r0

COPY ${checks} ./

HEALTHCHECK --interval=10s --retries=1 \
    CMD ./check_health

ENV INTERVAL=30

CMD ./compute_health $CHECKS
