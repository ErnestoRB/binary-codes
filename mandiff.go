package main

func mandiff(b uint8) {
	var prevVal uint8 = 0
	var transition= false
	arr := toBinaryArray(b)
	for i := 0; i < 8; i++ {
		nBit := arr[i]
		drawMandiffBit(nBit, prevVal, &transition, i)
		prevVal = nBit
	}
}

func drawMandiffBit(b, prevVal uint8, transition *bool, pos int) {
	if b == 0 {
		if !(*transition) {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0,1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0.5,0})
		} else {
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0,1})
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0,-1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0.5,0})
		}
	} else if b == 1 {
		*transition = !(*transition)
		if !(*transition){
			drawLine(Dc, Vector{ X: float64(pos), Y: 0}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0,1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0.5,0})
		} else {
			drawLine(Dc, Vector{ X: float64(pos), Y: 1}, Vector{0.5,0})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 1}, Vector{0,-1})
			drawLine(Dc, Vector{ X: (float64(pos) + 0.5), Y: 0}, Vector{0.5,0})
		}
		
	}
}