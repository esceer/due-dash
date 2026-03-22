#!make
include .local-env
export $(shell sed 's/=.*//' .local-env)

run:
	go run cmd/main.go
