FROM golang:1.12.7-alpine3.10 AS build

COPY src/*.go ./

RUN go build compute_health.go shared.go \
 && go build check_health.go shared.go

# # # # # # # # # # # # # # # # # # # # #

FROM healthcheck AS run

COPY --from=build \
    /go/compute_health \
    /go/check_health \
    ./
