.PHONY: install start-cli start-ui clean-logs lint

OS := $(shell uname -s)

install:
	go mod tidy

start-cli:
ifeq ($(OS),Linux)
	go run -ldflags="-extldflags=-Wl,--no-warn-duplicate-libraries" main.go input/*.torrent
else
	go run -ldflags="-extldflags -Wl,-no_warn_duplicate_libraries" main.go input/*.torrent
endif

start-ui:
ifeq ($(OS),Linux)
	go run -ldflags="-extldflags=-Wl,--no-warn-duplicate-libraries" main.go -ui
else
	go run -ldflags="-extldflags -Wl,-no_warn_duplicate_libraries" main.go -ui
endif

clean-logs:
	rm -rf logs
	mkdir logs

lint:
	golangci-lint run
