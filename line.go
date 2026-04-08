package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type line struct {
	label  string
	points plotter.XYs
}

func CreateLineplotPlot(title string, labels labels, bounds bounds, file string, points ...plotter.XYs) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = labels.x
	p.Y.Label.Text = labels.y
	p.X.Min, p.X.Max, p.Y.Min, p.Y.Max = bounds.xmin, bounds.xmax, bounds.ymin, bounds.ymax

	for _, pts := range points {
		err := plotutil.AddLines(p, "", pts)
		if err != nil {
			log.Fatalf("could not create lineplot: %+v", err)
		}
	}

	err := p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save lineplot: %+v", err)
	}
}
