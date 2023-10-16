include .env
export

CURRENT_DIR=$(shell pwd)

.PHONY: generate-gql

generate-gql:
	@go get github.com/99designs/gqlgen &&\
	go run github.com/99designs/gqlgen generate