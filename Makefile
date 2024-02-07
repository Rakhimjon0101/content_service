CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=${REGISTRY}
TAG=${TAG}
TAG_LATEST=${TAG_LATEST}
PKG_LIST := $(shell go list ${CURRENT_DIR}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
YAML_FILES :=  $(shell find .  -iname *-ci.yml  | grep -v /vendor/)

build: ## build
	CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-init: ## init swagger
	swag init -g cmd/main.go -o api/docs

run: ## run application
	go run cmd/main.go

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

lint-go: ## Lint go the files
	@golangci-lint run ./... || staticcheck -tests=false ./...

migrate-up:
	migrate -path migrations -database "postgres://admin:147ajt369@localhost:5555/dbo_content_service?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://admin:147ajt369@localhost:5555/dbo_content_service?sslmode=disable" down
