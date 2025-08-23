up:
	docker-compose up -d --no-deps --build

up-bff:
	docker compose up --build pmb-service pmb-db pmb-adminer -d

run:
	go run cmd/playcount-monitor/main.go || exit 1

run-tracker:
	go run cmd/playcount-tracker/main.go || exit 1

migrate:
	sql-migrate up -env="local"

deps:
	go mod tidy; go mod vendor

build:
	go build -o bin/bff cmd/playcount-monitor/main.go

build-tracker:
	go build -o bin/tracker cmd/playcount-tracker/main.go

build-cleaner:
	go build -o bin/cleaner cmd/db-cleaner/main.go

build-all:
	go build -o bin/bff cmd/playcount-monitor/main.go
	go build -o bin/tracker cmd/playcount-tracker/main.go
	go build -o bin/cleaner cmd/db-cleaner/main.go

test:
	go test ./... -v -cover

integration-test:
    ifeq ($(OS),Windows_NT)
	    set "RUN_INTEGRATION_TEST=true" && go test ./tests/... -v -cover
    else
	    export RUN_INTEGRATION_TEST=true && go test ./tests/... -v -cover
    endif

lint:
	golangci-lint run

add-sample-data:
	go run tests/data/add_sample_data.go

