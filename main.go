package main

import (
	"os"
	"strconv"

	"github.com/fogleman/gg"
)
var args = os.Args[1:]

const(
    s = 1000
    xGridNumber = 9
    yGridNumber = 9
) 

var widthY = float64(s / xGridNumber) // anchos de la celda
var widthX = float64(s / yGridNumber)

type Vector struct {
	X, Y float64
}

var origin = Vector{X: 0, Y: -yGridNumber/ 2}

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
	error := Dc.LoadFontFace("fonts/UbuntuMono.ttf",100)
	if(error != nil) {
		println("No se pudo cargar correctamente la fuente")
		os.Exit(1)
	}
	byteValue := uint8(value)
    setup(Dc, byteValue, "NRZL")
	nrzl(byteValue)
	Dc.SavePNG("nrzl.png")
    setup(Dc, byteValue, "NRZI")
	nrzi(byteValue)
	Dc.SavePNG("nrzi.png")
    setup(Dc, byteValue, "Bipolar AMI")
	bami(byteValue)
	Dc.SavePNG("bipolar_ami.png")
    setup(Dc, byteValue, "Pseudoternario")
	ptern(byteValue)
	Dc.SavePNG("pseudo_tern.png")
	setup(Dc, byteValue, "Manchester")
	mnchstr(byteValue)
	Dc.SavePNG("manchester.png")
	setup(Dc, byteValue, "Manchester Diferencial")
	mandiff(byteValue)
	Dc.SavePNG("man_differential.png")
}

func drawLine(dc *gg.Context, base, dir  Vector) {
	displacedBase := changeOrigin(base)
	normalizedBase := normalizeVector(displacedBase)
	final := Vector{displacedBase.X + dir.X, displacedBase.Y + dir.Y}
	normalizedFinal := normalizeVector(final)
	dc.DrawLine(normalizedBase.X, -normalizedBase.Y , normalizedFinal.X, -normalizedFinal.Y)
	dc.SetRGB(0,0,1)
	dc.SetLineWidth(10)
	dc.Stroke()
}

func drawTitle(dc *gg.Context, title string) {
	dc.SetRGB(1,0,0)
	dc.DrawString(title,10,100)
}

func setup(dc *gg.Context, b uint8, title string) {
	dc.SetRGB(1,1,1)
	dc.Clear()
	drawTitle(dc,title)
	dc.SetRGB(1,0,0)
	dc.DrawString(binaryString(b),10,200)
	drawTextAt(dc, Vector{8,0}, "0")
	drawTextAt(dc, Vector{8,1}, "1")
	drawTextAt(dc, Vector{8,-1}, "-1")
	arr := toBinaryArray(b)
	for pos := 0; pos < 8; pos++ {
		drawTextAt(Dc,Vector{float64(pos),-2}, strconv.FormatUint(uint64(arr[pos]), 2))
	}
	grid(dc)
}

func drawTextAt(dc *gg.Context, base  Vector, str string) {
	dc.SetRGB(0,1,0)
	realBase := normalizeVector(changeOrigin(base))
	dc.DrawString(str,realBase.X,-realBase.Y)
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
