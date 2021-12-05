package shuffler

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/spf13/cobra"
	"image"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func Shuffle(cmd *cobra.Command, args []string) {
	fyneApp := app.New()
	window := fyneApp.NewWindow("picture-shuffler")
	window.SetFullScreen(true)
	text := canvas.NewText("images coming soon", image.Black)
	window.SetContent(text)

	go changeContent(window, args[0])

	window.ShowAndRun()
}

func changeContent(fyneWindow fyne.Window, path string) {
	var err error
	for _ = 1; true; {
		err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			fmt.Println(path)
			openFile, err := os.Open(path)
			if err != nil {
				fmt.Println("some error on os.open***" + err.Error())
				return nil
			}
			defer func(openFile *os.File) {
				_ = openFile.Close()
			}(openFile)

			imageResult, _, err := image.Decode(openFile)
			if err != nil {
				fmt.Println("some error on image.decode***" + err.Error())
				return nil
			}

			img := canvas.NewImageFromImage(imageResult)
			fyneWindow.SetContent(img)
			time.Sleep(10 * time.Second)

			return nil
		})
	}

	if err != nil {
		return
	}
	fyneWindow.Close()
}
