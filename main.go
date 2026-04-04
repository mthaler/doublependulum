package main

import (
	"gonum.org/v1/plot/plotter"
)

const (
	x0     = 0.3
	phimax = 0.5
	g      = 9.81
	l      = 4.0
)

func main() {
	points := plotter.XYs{}
	for i := 0; i <= 500; i++ {
		t := float64(i) / 100.0
		points = append(points, plotter.XY{
			X: t,
			Y: x(t),
		})
	}
	b := bounds{
		xmin: 0,
		xmax: 5,
		ymin: -0.5,
		ymax: 0.5,
	}
	l := labels{
		x: "t",
		y: "x",
	}
	CreateLineplotPlot(points, "t - x", l, b, "eom.png")
	points2 := plotter.XYs{}
	for i := 0; i <= 500; i++ {
		t := float64(i) / 100.0
		points2 = append(points2, plotter.XY{
			X: x(t),
			Y: xdot(t),
		})
	}
	b2 := bounds{
		xmin: -0.3,
		xmax: 0.3,
		ymin: -0.5,
		ymax: 0.5,
	}
	l2 := labels{
		x: "x",
		y: "xdot",
	}
	CreateLineplotPlot(points2, "x - xdot", l2, b2, "pd.png")
}
