FROM alpine:3.11 AS build

RUN apk update && \
    apk add --no-cache \
        crystal=0.31.1-r2 \
        shards=0.9.0-r1 \
        libc-dev=0.7.2-r0

WORKDIR crystal
COPY src/*.cr ./

RUN crystal build compute_health.cr --release --static \
 && crystal build check_health.cr --release --static

# # # # # # # # # # # # # # # # # # # # #

FROM healthcheck AS run

COPY --from=build \
    /crystal/compute_health \
    /crystal/check_health \
    ./
