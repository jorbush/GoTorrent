.PHONY: install start-cli start-ui clean-logs debian-test linter

install:
	go mod tidy

start-cli:
	go run *.go input/*.torrent

start-ui:
	go run -ldflags="-extldflags=-Wl,-no_warn_duplicate_libraries" main.go -ui

clean-logs:
	rm -rf logs
	mkdir logs

lint:
	golangci-lint run
