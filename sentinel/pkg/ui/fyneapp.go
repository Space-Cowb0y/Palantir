package ui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunFyne() {
	a := app.New()
	w := a.NewWindow("Sentinel")

	openWeb := widget.NewButton("Open Web UI", func() {
		u, _ := url.Parse("http://localhost:8080/ui")
		_ = a.OpenURL(u)
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Sentinel Desktop GUI"),
		openWeb,
		widget.NewButton("Quit", func() { w.Close() }),
	))

	w.Resize(fyne.NewSize(420, 260))
	w.ShowAndRun()
}
