package main

func nrzi(b uint8) {
	var wasHigh = false
	arr := toBinaryArray(b)
	for i := 0; i < 8; i++ {
		nBit := arr[i]
		drawNrziBit(nBit, &wasHigh, i)
	}
}

func drawNrziBit(b uint8, wasHigh *bool, pos int) {
	if b == 1 {
		drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
		if(*wasHigh) {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{1,0})
		} else {
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{1,0})
		}
		*wasHigh = !(*wasHigh)
	} else if b == 0 {
		if(*wasHigh) {
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{1,0})
		} else {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{1,0})
		}
	}
}