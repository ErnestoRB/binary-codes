package main

import "github.com/fogleman/gg"

const(
    s = 1000
    xGridNumber = 10
    yGridNumber = 10
) 

var offsetY = float64(s / xGridNumber)
var offsetX = float64(s / yGridNumber)


func main() {
	
	dc := gg.NewContext(s, s)
    setup(dc)
	dc.SavePNG("out2.png")
}


func setup(dc *gg.Context) {
    dc.DrawRectangle(0, 0, s, s)
	dc.Fill()
	grid(dc)
}

func grid(dc *gg.Context) {
	for i := 1; i < xGridNumber; i++ {
        y := float64( i* int(offsetY))
		dc.DrawLine(0, y, s, y)
		dc.SetLineWidth(5)
		dc.SetRGB(0, 0, 0)
        dc.Stroke()
	}
    for i := 1; i < yGridNumber; i++ {
        x := float64( i* int(offsetX))
		dc.DrawLine(x, 0, x, s)
		dc.SetLineWidth(5)
		dc.SetRGB(0, 0, 0)
        dc.Stroke()
	}
}
