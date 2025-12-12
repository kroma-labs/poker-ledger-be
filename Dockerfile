# Build Stage
FROM golang:1.25 AS build-stage

ARG VERSION
ARG GIT_COMMIT
ARG GIT_BRANCH

LABEL app="build-poker-ledger-be"
LABEL REPO="https://github.com/kroma-labs/poker-ledger-be"

ENV PROJPATH=/app

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /app
COPY . .

RUN make build-alpine

# Final Stage
FROM alpine:latest

LABEL REPO="https://github.com/kroma-labs/poker-ledger-be"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Add tz data
RUN apk add --no-cache tzdata curl jq openssl
RUN apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing grpcurl

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/app
ENV TZ=UTC

WORKDIR /app

COPY --from=build-stage /app/.env.example .env
COPY --from=build-stage /app/bin/poker-ledger-be .
RUN chmod +x poker-ledger-be

# Create appuser
RUN adduser -D -g '' poker-ledger-be
USER poker-ledger-be
