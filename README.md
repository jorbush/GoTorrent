# GoTorrent

A BitTorrent client implementation in Go, designed to download files efficiently using the BitTorrent protocol.

## Requirements

- [Go](https://golang.org/doc/install)

## Installation

Run the following command to install the project dependencies:

```bash
make install
```

## Usage UI

### Run the App

Run the client using the make command:

```bash
make start-ui
```

You can drag and drop a torrent file into the app to start downloading the content or select a file using the `Browse` button.

Once the download is complete, the downloaded file will be saved in the `output` directory.

## Usage CLI

### Download a Torrent File

Download a torrent file, for example, the Debian installation image:

```bash
curl -o input/debian.torrent https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/debian-12.5.0-amd64-netinst.iso.torrent
```
### Run CLI

Run the client using the make command:

```bash
make start-cli
```

This command will start the BitTorrent client, which will download the content of the torrent file saved in the `input` directory. For example, the Debian image.

### Output

The downloaded file will be saved in the output directory.

## Linting

Check the code for style issues:
```bash
make lint
```

## References

- [BitTorrent Article](https://blog.jse.li/posts/torrent/)
- [The BitTorrent Protocol Specification](https://www.bittorrent.org/beps/bep_0003.html)
