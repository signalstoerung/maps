package maps

import (
	"fmt"
)

/*
MercatorAsciiPlot is a struct that holds the information to generate an ASCII plot of points.
Do not create one directly but instead use NewAsciiPlot(width, height)
*/
type MercatorAsciiPlot struct {
	height     uint
	width      uint
	plotbuffer []byte
}

// NewAsciiPlot returns a pointer to an MercatorAsciiPlot with the dimensions width * height.
func NewAsciiPlot(width uint, height uint) *MercatorAsciiPlot {
	pb := make([]byte, width*height)
	for i := range pb {
		pb[i] = ' '
	}
	// draw zero meridian
	for i := 0; i < int(height); i++ {
		pb[uint(i)*width+width/2] = '|'
	}
	for i := 0; i < int(width); i++ {
		location := height/2*width + uint(i)
		if i == int(width)/2 {
			pb[location] = '+'
		} else {
			pb[location] = '-'
		}
	}

	plot := MercatorAsciiPlot{width: width, height: height, plotbuffer: pb}
	return &plot
}

// AddPoint adds the supplied symbol to the plot.
// The maximum for x is width-1, for y height-1
func (a *MercatorAsciiPlot) AddPoint(x uint, y uint, symbol byte) error {
	// x = column, y = row
	err := fmt.Errorf("Invalid point.")
	switch {
	case x > a.width-1:
		return err
	case y > a.height-1:
		return err
	}
	location := y*a.width + x
	a.plotbuffer[location] = symbol
	return nil
}

// func (a *AsciiPlot) Read(p []byte) (n int, err error) {
// }

// DEBUG ONLY. Print is a hacked-together function to print the output; this will be replaced by a function
// implementing the io.Reader interface.
func (a MercatorAsciiPlot) Print() {
	for i := a.height - 1; i > 0; i-- {
		start := uint(i) * a.width
		end := (uint(i) + 1) * a.width

		line := string(a.plotbuffer[start:end])
		//		fmt.Printf("%x\n",line)
		fmt.Println(line)
	}
}
