package main

import (
	"fmt"
	"log"
	"os"
	"p2p/bittorrent/core"
	"p2p/ui"
	"time"
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

	startTime := time.Now()
	log.Printf("Timer started.\n")

	err = tf.DownloadTorrent()
	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(startTime)

	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60

	fmt.Printf("Download completed in %d minutes and %d seconds.\n", minutes, seconds)
	log.Printf("Downloaded in %v \n", duration)
}
