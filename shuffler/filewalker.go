package shuffler

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image"
	"io/fs"
	"os"
	"time"
)

func FileWalkerFunction(fyneWindow fyne.Window, path string, d fs.DirEntry, err error) error {
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
}
