set dotenv-load

app_name := "template"

[private]
default: help

help:
    just --list --justfile {{justfile()}}

init:
    -cp --no-clobber resources/config.yaml.example resources/config.yaml

##########
# Go
##########

fmt:
    gofmt -s -w .
    golangci-lint run --fix --config=.golangci.yaml

lint:
    golangci-lint run --config=.golangci.yaml

# TODO Fix lazy variable assignment https://github.com/casey/just/issues/953
kill:
    #!/usr/bin/env bash
    set -euxo pipefail
    server_port=$(yq '.server.port' resources/config.yaml)
    pid=$(lsof -t -i :$server_port) || true
    if [ "$pid" ]; then
      kill -9 $pid || true
    fi

run: #kill
    go run ./cmd/api

clean:
    go clean
    rm -rf ./out

tidy:
    go mod tidy

update:
    go get -u ./...
    just tidy

test:
    - just db-test-up
    - just db-test-migrate
    - go test -v ./...
    - just db-test-down

##########
# dbmate
##########
main_db_migrations := "resources/db/migrations"
main_db_schema := "resources/db/schema.sql"

db action:
    dbmate -u postgres://developer:secretpassword@localhost:5432/hotel_service?sslmode=disable -d {{main_db_migrations}} -s {{main_db_schema}} {{action}}

new-migration name:
    dbmate -d {{ main_db_migrations }} new {{ name }}

db-test-migrate:
    dbmate -u postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -d {{main_db_migrations}} -s {{main_db_schema}} up

##########################
# Infrastructure setup
##########################
db-setup:
    chmod u+x deployments/postgres.sh
    ./deployments/postgres.sh

db-test-up:
    docker container run --name test-postgres \
        -e POSTGRES_PASSWORD=postgres \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_DB=postgres \
        -p 5432:5432 \
        -d postgres:17.5
    sleep 2

db-test-down:
    docker container stop test-postgres && docker container rm test-postgres