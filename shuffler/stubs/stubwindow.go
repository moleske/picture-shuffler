package stubs

import "fyne.io/fyne/v2"

type StubWindow struct {
}

func (w *StubWindow) Canvas() fyne.Canvas {
	return nil
}

func (w *StubWindow) CenterOnScreen() {
}

func (w *StubWindow) Clipboard() fyne.Clipboard {
	return nil
}

func (w *StubWindow) Close() {
}

func (w *StubWindow) Content() fyne.CanvasObject {
	return nil
}

func (w *StubWindow) FixedSize() bool {
	return false
}

func (w *StubWindow) FullScreen() bool {
	return false
}

func (w *StubWindow) Hide() {
}

func (w *StubWindow) Icon() fyne.Resource {
	return nil
}

func (w *StubWindow) MainMenu() *fyne.MainMenu {
	return nil
}

func (w *StubWindow) Padded() bool {
	return false
}

func (w *StubWindow) RequestFocus() {
}

func (w *StubWindow) Resize(size fyne.Size) {
}

func (w *StubWindow) SetContent(obj fyne.CanvasObject) {
}

func (w *StubWindow) SetFixedSize(fixed bool) {
}

func (w *StubWindow) SetIcon(_ fyne.Resource) {
}

func (w *StubWindow) SetFullScreen(fullScreen bool) {
}

func (w *StubWindow) SetMainMenu(menu *fyne.MainMenu) {
}

func (w *StubWindow) SetMaster() {
}

func (w *StubWindow) SetOnClosed(closed func()) {
}

func (w *StubWindow) SetCloseIntercept(callback func()) {
}

func (w *StubWindow) SetPadded(padded bool) {
}

func (w *StubWindow) SetTitle(title string) {
}

func (w *StubWindow) Show() {
}

func (w *StubWindow) ShowAndRun() {
}

func (w *StubWindow) Title() string {
	return "Title of Window"
}
