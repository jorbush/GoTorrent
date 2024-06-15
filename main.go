package main

import (
	"flag"
	"fmt"
	"gotorrent/bittorrent/core"
	"gotorrent/ui"
	"image/color"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
	isDownloading := false
	myApp := app.New()
	myWindow := myApp.NewWindow("GoTorrent")
	myWindow.Resize(fyne.NewSize(600, 400))

	label := widget.NewLabel("Drag and drop a torrent file or click to browse")
	fileEntry := widget.NewEntry()
	fileEntry.SetPlaceHolder("No file selected")

	fileButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				log.Println("Failed to open file:", err)
				return
			}
			if reader == nil {
				dialog.ShowInformation("No file selected", "Please select a file", myWindow)
				log.Println("No file selected")
				return
			}
			fileEntry.SetText(reader.URI().Path())
		}, myWindow)
		fileDialog.Show()
	})

	startButton := widget.NewButton("Start Download", func() {
		if isDownloading {
			dialog.ShowInformation("Download in Progress", "A download is already in progress", myWindow)
			return
		}

		filePath := fileEntry.Text
		if filePath == "" {
			dialog.ShowInformation("No file selected", "Please select a file to download", myWindow)
			log.Println("No file selected for download")
			return
		}

		isDownloading = true
		go func() {
			defer func() { isDownloading = false }()
			startDownload(filePath, myWindow)
		}()
	})

	dropArea := canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	dropArea.SetMinSize(fyne.NewSize(400, 200))
	dropLabel := widget.NewLabel("Drop .torrent file here")

	dropContainer := container.NewStack(dropArea, container.NewCenter(dropLabel))

	myWindow.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		if len(uris) > 0 {
			fileEntry.SetText(uris[0].Path())
			dialog.ShowInformation("File Selected", fmt.Sprintf("Selected file: %s", uris[0].Path()), myWindow)
		}
	})

	content := container.NewVBox(label, fileEntry, dropContainer, fileButton, startButton)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func startCLI() {
	inPath := os.Args[1]
	startDownload(inPath, nil)
}

func startDownload(filePath string, myWindow fyne.Window) {
	tf, err := core.Open(filePath)
	if err != nil {
		if myWindow != nil {
			dialog.ShowError(err, myWindow)
		}
		log.Fatal(err)
	}

	startTime := time.Now()
	log.Printf("Timer started.\n")

	err = tf.DownloadTorrent()
	if err != nil {
		if myWindow != nil {
			dialog.ShowError(err, myWindow)
		}
		log.Fatal(err)
	}

	duration := time.Since(startTime)
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60

	message := fmt.Sprintf("Download completed in %d minutes and %d seconds.\n", minutes, seconds)
	if myWindow != nil {
		dialog.ShowInformation("Download Completed", message, myWindow)
	}
	log.Printf("Downloaded in %v \n", duration)
}
