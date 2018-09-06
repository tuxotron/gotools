package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) < 1 {
		panic("Missing file")
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	check(err)

	freqList := make([]float64, 256)
	size := float64(len(bytes))

	for i := range [256]int{} {

		cont := 0
		for _, v := range bytes {
			if i == int(v) {
				cont++
			}
		}
		freqList[i] = float64(cont)/size
	}

	ent := 0.0
	for _, f := range freqList {
		if f > 0.0 {
			ent = ent + (f * math.Log2(f))
		}
	}

	fmt.Println("Shannon entropy: ", -ent)

	p, err := plot.New()
	check(err)

	p.Title.Text = os.Args[1] + " - Shannon Entropy: " + fmt.Sprintf("%f", -ent)
	p.X.Label.Text = "byte"
	p.Y.Label.Text = "Freq"
	p.Y.Min = 0
	p.Y.Max = 1
	p.X.Min = 0
	p.X.Max = 255

	err = plotutil.AddLinePoints(p, "", getPoints(freqList))
	check(err)

	// Save the plot to a PNG file.
	err = p.Save(8*vg.Inch, 4*vg.Inch, os.Args[1] + "_plot.png")
	check(err)
}

func getPoints(entropyPts []float64) plotter.XYs {

	pts := make(plotter.XYs, len(entropyPts))

	for i, v := range entropyPts {
		pts[i].Y = v
		pts[i].X = float64(i)
	}

	return pts
}
