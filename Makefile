.PHONY: install run clean-logs

install:
	go mod tidy

run:
	go run *.go input/debian.torrent output/debian.iso

clean-logs:
	rm -rf logs
	mkdir logs
