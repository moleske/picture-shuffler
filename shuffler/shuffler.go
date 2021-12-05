package shuffler

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/spf13/cobra"
	"image"
	"io/fs"
	"path/filepath"
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
			return FileWalkerFunction(fyneWindow, path, d, err)
		})
	}

	if err != nil {
		return
	}
	fyneWindow.Close()
}
