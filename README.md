# P2P Client

Implementing a [BitTorrent](https://www.bittorrent.org/beps/bep_0003.html) client.

## Requirements

- [Go](https://golang.org/doc/install)

## Usage

Download the torrent image, for example for debian:

```bash
curl -o input/debian.torrent https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/debian-12.5.0-amd64-netinst.iso.torrent
```

Run the client:

```bash
make run
```

This will download the content of the torrent file saved in `input` directory. In this case, the debian image.

You will found the downloaded file in the `output` directory.

## References

- [BitTorrent Article](https://blog.jse.li/posts/torrent/)
- [The BitTorrent Protocol Specification](https://www.bittorrent.org/beps/bep_0003.html)
