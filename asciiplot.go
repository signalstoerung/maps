package maps

import (
	"fmt"
)

const (
	TerminalScale = 80 / (2*math.Pi)
)

type AsciiPlot struct {
	height uint
	width uint
	plotbuffer []byte
}

func (a *AsciiPlot) Init(width uint, height uint) {
	a.width=width
	a.height=height
	a.plotbuffer=make([]byte,0,width*height)
}

func (a *AsciiPlot) AddPoint(x uint, y uint) {

}