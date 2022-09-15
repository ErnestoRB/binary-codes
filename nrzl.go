package main

func nrzl(b uint8) {
	var prevVal uint8
	arr := toBinaryArray(b)
	for i := 0; i < 8; i++ {
		nBit := arr[i]
		drawNrzlBit(nBit, prevVal, i)
		prevVal = nBit;
	}
}

func drawNrzlBit(b, prevVal uint8, pos int) {
	if b == 1 {
		drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{1,0})
		if prevVal == 0 {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
		}
	} else if b == 0 {
		drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{1,0})
		if prevVal == 1 {
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{0,-1})
		}
	}
}