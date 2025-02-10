set dotenv-load

setup:
	go install github.com/bitfield/gotestdox/cmd/gotestdox@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

format:
	gofumpt -l -w .
	goimports-reviser -rm-unused -set-alias ./...
	golines -w -m 120 .

# build -> build application
build:
	go build -o ./dist/main ./cmd

# run -> application
run:
	./dist/main

# dev -> run build then run it
dev: 
	watchexec -r -c -e go -- just build run

# test -> testing
test:
  gotestdox -v ./...

# health -> Hit Health Check Endpoint
health:
	curl -s http://localhost:8000/healthz | jq

# migrate-create -> create migration
migrate-create NAME:
	migrate create -ext sql -dir ./database/migrations -seq {{NAME}}

# migrate-up -> up migration
migrate-up:
	migrate -path ./database/migrations -database "$DATABASE_URL" up

# migrate-down -> down migration
migrate-down:
	migrate -path ./database/migrations -database "$DATABASE_URL" down