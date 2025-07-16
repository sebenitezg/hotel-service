set dotenv-load

app_name := "template"

[private]
default: help

help:
    just --list --justfile {{justfile()}}


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

##########
# dbmate
##########
main_db_migrations := "resources/db/migrations"
main_db_schema := "resources/db/schema.sql"

db action:
    dbmate -u postgres://developer:123456@localhost:5432/postgres?sslmode=disable -d {{main_db_migrations}} -s {{main_db_schema}} {{action}}

new-migration name:
    dbmate -d {{ main_db_migrations }} new {{ name }}
