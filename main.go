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
	outPath := os.Args[2]

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

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
