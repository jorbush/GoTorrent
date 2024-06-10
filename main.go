package main

import (
	"fmt"
	"log"
	"os"
	"p2p/bittorrent/core"
	"p2p/ui"
)

func main() {
	inPath := os.Args[1]

	logFile, err := ui.SetupLogger()
	if err != nil {
		fmt.Printf("Error setting up logger: %v\n", err)
		return
	}
	defer logFile.Close()

	tf, err := core.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadTorrent()
	if err != nil {
		log.Fatal(err)
	}
}
