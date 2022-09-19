package main

var up int

func mnchstr(b uint8) {
	var prevVal uint8 = 0
	var transition= false
	arr := toBinaryArray(b)
	for i := 0; i < 8; i++ {
		nBit := arr[i]
		drawMnchstrBit(nBit, prevVal, &transition, i)
		prevVal = nBit
	}
}

func drawMnchstrBit(b, prevVal uint8, transition *bool, pos int)  {
	if pos == 0 {
		if b == 1 {
			up = 0;
		} else{
			up = 1;
		}
	}

	if b == 1 {
		if up == 1 {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0,1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0.5,0})
		} else {
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0,-1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0.5,0})
		}
	}

	if b == 0 {
		if up == 0 {
			if prevVal == 0 {
				drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
			}
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0,1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0.5,0})
		} else {
			if prevVal == 0 {
				drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
			}
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0,-1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0.5,0})
		}
	}
}