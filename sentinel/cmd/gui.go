package cmd


import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func RunGUI() {
	a := app.New()
	w := a.NewWindow("Sentinel GUI")


	label := widget.NewLabel("Welcome to Sentinel GUI")
	btn := widget.NewButton("Ping Agent", func() {
		label.SetText("Ping sent!")
	})


	w.SetContent(container.NewVBox(label, btn))
	w.ShowAndRun()
}