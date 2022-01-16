# SRC := $(shell find . -type f -name "*.go")
BIN := ./gochan

all:
	go build -o $(BIN) main.go
