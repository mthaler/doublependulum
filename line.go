package main

import (
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func CreateLineplotPlot(title string, labels labels, bounds bounds, file string, lines ...*plotter.Line) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = labels.x
	p.Y.Label.Text = labels.y
	p.X.Min, p.X.Max, p.Y.Min, p.Y.Max = bounds.xmin, bounds.xmax, bounds.ymin, bounds.ymax
	p.BackgroundColor = color.Transparent

	for _, line := range lines {
		err := plotutil.AddLines(p, line)
		if err != nil {
			log.Fatalf("could not create lineplot: %+v", err)
		}
	}

	c := vgimg.PngCanvas{vgimg.NewWith(
		vgimg.UseWH(10*vg.Centimeter, 5*vg.Centimeter),
		vgimg.UseBackgroundColor(color.Transparent),
	)}
	p.Draw(draw.New(c))

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = c.WriteTo(f)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
