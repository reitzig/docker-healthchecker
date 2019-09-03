FROM ruby:2.6.5-alpine3.11 AS prepare

RUN apk update \
 && apk add --no-cache  \
    build-base=0.5-r1  \
    ruby-dev=2.6.5-r2 \
 && gem install bundler

WORKDIR app
COPY Gemfile ./
RUN bundle install --path=gems

# # # # # # # # # # # # # # # # # # # # #

FROM healthcheck AS run

RUN apk update \
 && apk add --no-cache  \
    ruby=2.6.5-r2 \
 && gem install --no-document bundler

# Copy gems over from build stage
WORKDIR app
COPY --from=prepare /app ./
COPY src/* ./
RUN bundle config --local path gems

WORKDIR /
COPY *_health ./
