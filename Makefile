.PHONY: install run

install:
	go mod tidy

run:
	go run *.go input/debian.torrent output/debian.iso
