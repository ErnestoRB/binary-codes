package main

func bami(b uint8) {
	var prevVal uint8 = 0
	var wasHigh = false
	arr := toBinaryArray(b)
	for i := 0; i < 8; i++ {
		nBit := arr[i]
		drawBamiBit(nBit, prevVal, &wasHigh, i)
		prevVal = nBit
	}
}

func drawBamiBit(b, prevVal uint8, wasHigh *bool, pos int) {
	if b == 1 {
		if(*wasHigh) {
			if prevVal == 0 {
				drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,-1})
				drawLine(Dc, Vector{ X: float64(pos), Y: -1}, Vector{1,0})
			} else {
				drawLine(Dc, Vector{ X: float64(pos), Y: -1}, Vector{0,2})
				drawLine(Dc, Vector{ X: float64(pos), Y: -1}, Vector{1,0})
			}
		} else {
			if prevVal == 0 {
				drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
				drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{1,0})
			} else {
				drawLine(Dc, Vector{ X: float64(pos), Y: -1}, Vector{0,2})
				drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{1,0})
			}
		}
		*wasHigh = !(*wasHigh)
	} else if b == 0 {
		drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{1,0})
		if(prevVal == 0) {
			return
		}
		if *wasHigh {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
		} else {
			drawLine(Dc, Vector{ X: float64(pos), Y: -1}, Vector{0,1})
		}
	}
}