up:
	docker-compose up --build osu-dashboard-bff -d
	docker-compose up --build osu-dashboard-jobs -d

up-bff:
	docker-compose up --build osu-dashboard-bff -d

up-jobs:
	docker-compose up --build osu-dashboard-jobs -d

run:
	go run cmd/bff/main.go || exit 1

run-jobs:
	go run cmd/jobs/main.go || exit 1

migrate:
	sql-migrate up -env="local"

deps:
	go mod tidy; go mod vendor

build:
	go build -o bin/bff cmd/bff/main.go

build-jobs:
	go build -o bin/jobs cmd/jobs/main.go

build-all:
	go build -o bin/bff cmd/bff/main.go
	go build -o bin/jobs cmd/jobs/main.go

test:
	go test ./... -v -cover

integration-test:
    ifeq ($(OS),Windows_NT)
	    set "RUN_INTEGRATION_TEST=true" && go test ./tests/... -v -cover
    else
	    export RUN_INTEGRATION_TEST=true && go test ./tests/... -v -cover
    endif

lint:
	golangci-lint run --fix --config=.golangci.yml

add-sample-data:
	go run tests/data/add_sample_data.go

