package main

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func CreateLineplotPlot(title string, labels labels, bounds bounds, file string, lines ...*plotter.Line) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = labels.x
	p.Y.Label.Text = labels.y
	p.X.Min, p.X.Max, p.Y.Min, p.Y.Max = bounds.xmin, bounds.xmax, bounds.ymin, bounds.ymax
	p.BackgroundColor = color.Transparent
	p.Add(plotter.NewGrid())

	for _, line := range lines {
		p.Add(line)
	}

	err := p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save lineplot: %+v", err)
	}
}
