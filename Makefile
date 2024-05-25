.PHONY: install run clean-logs debian-test

install:
	go mod tidy

debian-test:
	go run *.go input/debian.torrent output/debian.iso

clean-logs:
	rm -rf logs
	mkdir logs
