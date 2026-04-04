package main

import (
	"gonum.org/v1/plot/plotter"
)

const (
	x0     = 0.3
	phimax = 0.5
	g      = 9.81
	l1     = 3.0
	l2     = 3.0
)

func main() {
	points := plotter.XYs{}
	deltat := 1.0 / 100.0
	for i := 0; i <= 500; i++ {
		if i == 0 {
			// as inital value for x we use x0
			points = append(points, plotter.XY{
				X: 0.0,
				Y: x(x0, 0, deltat),
			})
		} else {
			t := float64(i) / 100.0
			xprev := points[i-1].X
			points = append(points, plotter.XY{
				X: t,
				Y: x(xprev, t, deltat),
			})
		}
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
}
