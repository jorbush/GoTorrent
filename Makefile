.PHONY: install run

install:
	go mod tidy

run:
	go run *.go debian-12.5.0-amd64-netinst.iso.torrent debian.iso
