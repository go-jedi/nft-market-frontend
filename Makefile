.PHONY: run

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

.DEFAULT_GOAL := run