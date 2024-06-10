.PHONY: install run clean-logs debian-test linter

install:
	go mod tidy

run:
	go run *.go input/*.torrent

clean-logs:
	rm -rf logs
	mkdir logs

linter:
	golangci-lint run
