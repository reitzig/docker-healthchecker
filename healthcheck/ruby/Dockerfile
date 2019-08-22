FROM ruby:2.5.5-alpine3.10 AS prepare

RUN apk update \
 && apk add --no-cache  \
    build-base=0.5-r1  \
    ruby-dev=2.5.5-r0 \
 && gem install bundler

WORKDIR app
COPY Gemfile ./
RUN bundle install --path=gems

# # # # # # # # # # # # # # # # # # # # #

FROM healthcheck AS run

RUN apk update \
 && apk add --no-cache  \
    ruby=2.5.5-r0 \
 && gem install --no-ri --no-rdoc bundler

# Copy gems over from build stage
WORKDIR app
COPY --from=prepare /app ./
COPY src/* ./
RUN bundle config --local path gems

WORKDIR /
COPY *_health ./
