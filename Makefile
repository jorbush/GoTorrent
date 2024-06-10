.PHONY: install run clean-logs debian-test linter

install:
	go mod tidy

debian-test:
	go run *.go input/*.torrent output/debian.iso

clean-logs:
	rm -rf logs
	mkdir logs

linter:
	golangci-lint run
