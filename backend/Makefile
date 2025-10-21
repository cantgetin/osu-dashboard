up:
	docker compose up --build pmb-service -d
	docker compose up --build pmb-worker -d
	docker compose up --build pmb-cleaner -d

up-bff:
	docker-compose up --build pmb-service -d

up-cleaner:
	docker-compose up --build pmb-cleaner -d

up-tracker:
	docker-compose up --build pmb-worker -d

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

