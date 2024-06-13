# P2P Client

A BitTorrent client implementation in Go, designed to download files efficiently using the BitTorrent protocol.

## Requirements

- [Go](https://golang.org/doc/install)

## Usage

### Download a Torrent File

Download a torrent file, for example, the Debian installation image:

```bash
curl -o input/debian.torrent https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/debian-12.5.0-amd64-netinst.iso.torrent
```
### Run the Client
Run the client using the make command:

```bash
make run
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
