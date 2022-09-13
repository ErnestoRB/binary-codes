package main

import (
	"os"
	"strconv"

	"github.com/fogleman/gg"
)
var args = os.Args[1:]

const(
    s = 1000
    xGridNumber = 10
    yGridNumber = 10
) 

var widthY = float64(s / xGridNumber) // anchos de la celda
var widthX = float64(s / yGridNumber)

type Vector struct {
	X, Y float64
}

var origin = Vector{X: 0, Y: yGridNumber/ 2}

var Dc = gg.NewContext(s, s)

func main() {
	if len(args) == 0 {
		println("Pasa un argumento numerico de un byte!")
		return
	}
	value, err := strconv.Atoi(args[0])
	if(err != nil || value > 255 || value < 0){
		println("No introduciste un número válido (byte) (0-255)")
		os.Exit(1)
	}
	byteValue := uint8(value)
    setup(Dc)
	nrzl(byteValue)
	Dc.SavePNG("nrzl.png")
    setup(Dc)
	nrzi(byteValue)
	Dc.SavePNG("nrzi.png")
    setup(Dc)
	bami(byteValue)
	Dc.SavePNG("bipolar_ami.png")
    setup(Dc)
	ptern(byteValue)
	Dc.SavePNG("pseudo_tern.png")
}

func drawLine(dc *gg.Context, base, dir  Vector) {
	realBase := Vector{X: (origin.X + base.X) * widthX, Y: (origin.Y - base.Y) * widthY}
	realDir := Vector{X: dir.X * widthX, Y: -dir.Y * widthY}
	dc.DrawLine(realBase.X, realBase.Y , realBase.X + realDir.X, realBase.Y + realDir.Y )
	dc.SetRGB(0,0,1)
	dc.SetLineWidth(10)
	dc.Stroke()
}

func setup(dc *gg.Context) {
	dc.SetRGB(1,1,1)
	dc.Clear()
	grid(dc)
}

func grid(dc *gg.Context) {
	dc.SetLineWidth(2)
	for i := 1; i < xGridNumber; i++ {
        y := float64( i* int(widthY))
		dc.DrawLine(0, y, s, y)
		dc.SetRGB(0, 0, 0)
        dc.Stroke()
	}
    for i := 1; i < yGridNumber; i++ {
        x := float64( i* int(widthX))
		dc.DrawLine(x, 0, x, s)
		dc.SetRGB(0, 0, 0)
        dc.Stroke()
	}
}
