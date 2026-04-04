package main

import (
	"fyne.io/fyne"
)

const (
	l1 = 10.0
	l2 = 20.0
)

func main() {
}

func show(app fyne.App) {
	window := app.NewWindow("Life")

	window.Resize(fyne.NewSize(600, 600))
	window.Show()
}
