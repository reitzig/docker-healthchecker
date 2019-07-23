FROM golang:1.12.7-alpine3.10 AS build

COPY src/*.go ./

RUN go build compute_health.go \
 && go build check_health.go

# # # # # # # # # # # # # # # # # # # # #

FROM alpine:3.10.1 AS run

ARG checks
ENV CHECKS=${checks}

RUN apk update && apk add --no-cache \
    postgresql-client=11.4-r0 \
    curl=7.65.1-r0

COPY --from=build \
    /go/compute_health \
    /go/check_health \
    ./

COPY ${checks} ./

HEALTHCHECK --interval=10s --retries=1 \
    CMD ./check_health

ENV INTERVAL=30

CMD ./compute_health $CHECKS
