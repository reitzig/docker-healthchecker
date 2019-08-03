FROM alpine:3.10.1 AS run

ARG checks
ENV CHECKS=checks.json \
    INTERVAL=30 \
    SUMMARY_DIR=/www \
    MINI_HTTP_CONF=/etc/mini_httpd/mini_httpd.conf

RUN apk update \
 && apk add --no-cache \
        postgresql-client=11.4-r0 \
        curl=7.65.1-r0 \
        mini_httpd=1.30-r0 \
 && mkdir -p "$SUMMARY_DIR" "$(dirname $MINI_HTTP_CONF)" \
 && echo -e "port=80\nuser=minihttpd\ndir=$SUMMARY_DIR\nchroot" \
        > "$MINI_HTTP_CONF" \
 && chown minihttpd "$SUMMARY_DIR" "$MINI_HTTP_CONF"

EXPOSE 80

COPY ${checks} ./$CHECKS

HEALTHCHECK --interval=10s --retries=1 \
    CMD ./check_health

CMD sh -c 'mini_httpd -C "$MINI_HTTP_CONF" -D &' \
 && ./compute_health checks.json
