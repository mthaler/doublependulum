package main

import (
	"math"

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

func calculatePos(theta1, theta2 float64) (x, y float64) {
	xx := l1*math.Sin(theta1) + l2*math.Sin(theta2)
	yy := l1*math.Cos(theta1) + l2*math.Cos(theta2)
	return xx, yy
}
