up:
	docker-compose up -d --no-deps --build

up-no-worker:
	docker-compose up --scale pmb-worker=0 -d --no-deps --build

run:
	go run cmd/playcount-monitor/main.go || exit 1

run-tracker:
	go run cmd/playcount-tracker/main.go || exit 1

migrate:
	sql-migrate up -env="local"

build:
	go build -o bin/main cmd/playcount-monitor/main.go

build-tracker:
	go build -o bin/main cmd/playcount-tracker/main.go

test:
	go test ./... -v -cover

integration-test:
    ifeq ($(OS),Windows_NT)
	    set "RUN_INTEGRATION_TEST=true" && go test ./tests/... -v -cover
    else
	    export RUN_INTEGRATION_TEST=true && go test ./tests/... -v -cover
    endif

add-sample-data:
	go run tests/data/add_sample_data.go

