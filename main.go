package main

import (
	"image/color"
	"log"
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
	for i := 0; i <= 4000; i++ {
		if i == 0 {
			// as inital value for x we use x0
			points = append(points, data{y1: th1 / math.Pi, y2: v1 / math.Pi, y3: th2 / math.Pi, y4: v2 / math.Pi})
		} else {
			deltat := 1 / 100.0
			t := float64(i) / 10.0
			d := angles(points[i-1], t, deltat)
			points = append(points, d)
		}
	}
	var points2 plotter.XYs
	for i := 0; i <= 4000; i++ {
		t := float64(i) / 10.0
		points2 = append(points2, plotter.XY{
			X: t,
			Y: l1*math.Sin(points[i].y1) + l2*math.Sin(points[i].y3),
		})
	}
	line1, err := plotter.NewLine(points2)
	if err != nil {
		log.Fatalf("could not create line: %+v", err)
	}
	//line1.LineStyle.Width = vg.Points(1)
	var points3 []data
	for i := 0; i <= 4000; i++ {
		if i == 0 {
			// as inital value for x we use x0
			points3 = append(points3, data{y1: (th1 + 0.01) / math.Pi, y2: v1 / math.Pi, y3: th2 / math.Pi, y4: v2 / math.Pi})
		} else {
			deltat := 1 / 100.0
			t := float64(i) / 10.0
			d := angles(points3[i-1], t, deltat)
			points3 = append(points3, d)
		}
	}
	var points4 plotter.XYs
	for i := 0; i <= 4000; i++ {
		t := float64(i) / 10.0
		points4 = append(points4, plotter.XY{
			X: t,
			Y: l1*math.Sin(points3[i].y1) + l2*math.Sin(points3[i].y3),
		})
	}
	line2, err := plotter.NewLine(points4)
	line2.LineStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	if err != nil {
		log.Fatalf("could not create line: %+v", err)
	}
	var points5 []data
	for i := 0; i <= 4000; i++ {
		if i == 0 {
			// as inital value for x we use x0
			points5 = append(points5, data{y1: (th1 + 0.02) / math.Pi, y2: v1 / math.Pi, y3: th2 / math.Pi, y4: v2 / math.Pi})
		} else {
			deltat := 1 / 100.0
			t := float64(i) / 10.0
			d := angles(points5[i-1], t, deltat)
			points5 = append(points5, d)
		}
	}
	var points6 plotter.XYs
	for i := 0; i <= 4000; i++ {
		t := float64(i) / 10.0
		points6 = append(points6, plotter.XY{
			X: t,
			Y: l1*math.Sin(points5[i].y1) + l2*math.Sin(points5[i].y3),
		})
	}
	line3, err := plotter.NewLine(points6)
	line3.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	if err != nil {
		log.Fatalf("could not create line: %+v", err)
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
	CreateLineplotPlot("t - x", l, b, "eom.png", line1, line2, line3)
}
