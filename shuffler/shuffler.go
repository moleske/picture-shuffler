package shuffler

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/spf13/cobra"
	"image"
	"image/color"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func Shuffle(cmd *cobra.Command, args []string) {
	fyneApp := app.New()
	window := fyneApp.NewWindow("picture-shuffler")
	window.SetFullScreen(true)
	window.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "Escape" {
			window.Close()
		}
	})
	go changeContent(window, args[0])

	window.ShowAndRun()
}

func changeContent(window fyne.Window, path string) {
	var err error
	var img *canvas.Image
	var canvasPathText *canvas.Text
	var c *fyne.Container
	var runonce = false
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

			imageResult, _, err := image.Decode(openFile)
			if err != nil {
				fmt.Println("some error on image.decode***" + err.Error())
				return nil
			}

			err = openFile.Close()
			if err != nil {
				fmt.Println("some error on file.close***" + err.Error())
			}

			if runonce == false {
				img = canvas.NewImageFromImage(imageResult)
				img.FillMode = canvas.ImageFillContain
				canvasPathText = canvas.NewText(path, color.White)
				c = container.New(layout.NewBorderLayout(canvasPathText, nil, nil, nil), canvasPathText, img)
				window.SetContent(c)
				runonce = true
			} else {
				*img = *canvas.NewImageFromImage(imageResult)
				*canvasPathText = *canvas.NewText(path, color.White)

				c.Refresh()
			}
			time.Sleep(10 * time.Second)

			return nil
		})
	}

	if err != nil {
		return
	}
}
