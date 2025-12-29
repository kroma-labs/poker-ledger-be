-include ${CURDIR}/.env

.PHONY: build clean test default install

BIN_NAME=poker-ledger

SHELL ?= $(shell which bash)
VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT ?= $(shell git rev-parse HEAD)
GIT_DIRTY ?= $(shell test -n "`git status --porcelain`" && echo "+DIRTY" || true)
BUILD_DATE ?= $(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := poker-ledger

default: test

help:
	@echo 'Management commands:'
	@echo
	@echo 'Usage:'
	@echo '    make bench           Run benchmarks.'
	@echo '    make build           Compile the project and generate a binary.'
	@echo '    make clean           Clean the directory tree.'
	@echo '    make coverage-html   Generate test coverage report.'
	@echo '    make swagger   		Generate swagger documentation'
	@echo '    make dep             Update dependencies.'
	@echo '    make help            Show this message.'
	@echo '    make lint            Run linters on the project.'
	@echo '    make mock            Generate mocks for interfaces.'
	@echo '    make run-http		Run HTTP server locally.'
	@echo '    make migrate-up      Apply database migrations.'
	@echo '    make migrate-down    Revert database migrations.'
	@echo '    make test            Run tests.'
	@echo

build:
	@echo "building ${BIN_NAME}"
	@echo "GOPATH=${GOPATH}"
	@CGO_ENABLED=0 GO111MODULE=on GOEXPERIMENT=greenteagc go build -tags="sonic avx no_clickhouse no_libsql no_sqlite3 no_mssql no_vertica no_mysql no_ydb" \
		-ldflags "-w -s" \
		-v -o bin/${BIN_NAME}

build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	@CGO_ENABLED=0 GO111MODULE=on GOEXPERIMENT=greenteagc go build -tags="sonic avx no_clickhouse no_libsql no_sqlite3 no_mssql no_vertica no_mysql no_ydb" \
		-ldflags "-w -s" \
		-v -o bin/${BIN_NAME}

build-docker:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	@docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) --build-arg GIT_BRANCH=$(GIT_BRANCH) -t $(IMAGE_NAME):local .

build-podman:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	@podman build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) --build-arg GIT_BRANCH=$(GIT_BRANCH) -t $(IMAGE_NAME):local .

dep:
	@go mod tidy

install:
	@go install github.com/vektra/mockery/v3@v3.5.0
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.5.0
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/pressly/goose/v3@latest

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}
	@go clean ./...

test:
	@go test -race -v \
		-coverpkg=$$(go list ./... | grep -v -E '/(mocks|cmd)($$|/)' | tr '\n' ',' | sed 's/,$$//') \
		-coverprofile=cover.out ./...
	@go tool cover -func=cover.out

coverage-html:
	@go test -race -v \
		-coverpkg=$$(go list ./... | grep -v -E '/(mocks|cmd)($$|/)' | tr '\n' ',' | sed 's/,$$//') \
		-coverprofile=cover.out ./...
	@go tool cover -html=cover.out && rm -rf cover.out

bench:
	# -run=^B negates all tests
	@go test -bench=. -run=^B -benchtime 10s -benchmem ./...

lint: install
	@golangci-lint fmt
	@golangci-lint run --fix

mock: install
	@mockery --config .mockery.yaml

swagger: install
	@swag init --generatedTime

run-http: build
	@echo
	@echo "swagger ui available at http://${HTTP_SERVER_HOST}:${HTTP_SERVER_PORT}/swagger/index.html"
	@echo
	@set -a; source .env; ./bin/${BIN_NAME} serve-http

migrate-up: build
	@echo
	@echo "applying database migrations"
	@echo
	@set -a; source .env; ./bin/${BIN_NAME} migrate-up

migrate-down: build
	@echo
	@echo "reverting database migrations"
	@echo
	@set -a; source .env; ./bin/${BIN_NAME} migrate-down

migrate-create: build
	@echo
	@echo "creating database migration"
	@echo
	@goose -dir db/migrations create ${name} sql                          

sqlite-dump:
	@echo
	@echo "dumping sqlite database"
	@echo
	@sqlite3 ./internal/adapters/db/sqlite/db.sqlite .schema > ./internal/adapters/db/sqlite/schema.sql

http-hot:
	@echo "🚀 Starting HTTP server with hot reload..."
	air --build.cmd "go build -o bin/http ./cmd/http" --build.bin "./bin/http"
