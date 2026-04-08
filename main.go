package main

import (
	"math"

	"gonum.org/v1/plot/plotter"
)

const (
	g   = 9.81 // Aceleration due to gravity
	l1  = 1.0  // Length of the first pendulum
	l2  = 1.0  // Length of the second pendulum
	m1  = 0.5  // Mass of the first pendulum
	m2  = 0.5  // Mass of the second pendulum
	th1 = 45   // Inital displacement of the first pendulum
	th2 = 0    // Inital displacement of the second pendulum
	v1  = 0    // initial velocity of the first pendulum
	v2  = 0    // initial velocity of the first pendulum
)

func main() {
	var points []data
	deltat := 1 / 100.0
	for i := 0; i <= 1000; i++ {
		if i == 0 {
			// as inital value for x we use x0
			points = append(points, data{y1: th1 / math.Pi, y2: v1 / math.Pi, y3: th2 / math.Pi, y4: v2 / math.Pi})
		} else {
			t := float64(i) / 10.0
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
