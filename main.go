package main

import (
	"flag"
	"fmt"
	"gotorrent/bittorrent/core"
	"gotorrent/ui"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	useUI := flag.Bool("ui", false, "Use graphical user interface")
	flag.Parse()

	logFile, err := ui.SetupLogger()
	if err != nil {
		fmt.Printf("Error setting up logger: %v\n", err)
		return
	}
	defer logFile.Close()

	if *useUI {
		startUI()
	} else {
		startCLI()
	}
}

func startUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GoTorrent")

	label := widget.NewLabel("Drag and drop a torrent file or click to browse")
	fileEntry := widget.NewEntry()
	fileEntry.SetPlaceHolder("No file selected")

	fileButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("Failed to open file:", err)
				return
			}
			if reader == nil {
				log.Println("No file selected")
				return
			}
			fileEntry.SetText(reader.URI().Path())
		}, myWindow)
		fileDialog.Show()
	})

	startButton := widget.NewButton("Start Download", func() {
		filePath := fileEntry.Text
		if filePath == "" {
			log.Println("No file selected for download")
			return
		}
		startDownload(filePath)
	})

	content := container.NewVBox(label, fileEntry, fileButton, startButton)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func startCLI() {
	inPath := os.Args[1]
	startDownload(inPath)
}

func startDownload(filePath string) {
	tf, err := core.Open(filePath)
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
