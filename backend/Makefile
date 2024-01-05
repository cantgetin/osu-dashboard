up:
	docker-compose up -d

run:
	go run cmd/main.go || exit 1

migrate:
	sql-migrate up -env="local"

build:
	go build -o bin/main cmd/main.go

test:
	go test ./tests/... -v -cover

add-sample-data:
	go run tests/data/add_sample_data.go

