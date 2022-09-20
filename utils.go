package main

import "strconv"

// Hace que el vector esté en terminos de ancho de celdas
func normalizeVector(base Vector) Vector {
	return Vector{X: base.X * widthX, Y: base.Y * widthY}
}

func changeOrigin(base Vector) Vector {
	return Vector{X: (origin.X + base.X), Y: (origin.Y + base.Y)}
}

func toBinaryArray(b uint8) (arr [8]uint8) {
	for i := 7; i >= 0; i-- {
		var maskBit uint8 = 1 << i
		nBit := b & maskBit
		nBit >>= i
		arr[7 - i] =nBit
	}
	return
}

func binaryString(b uint8) (str string) {
	str = strconv.FormatUint(uint64(b), 2)
	length := len(str) 
	if length < 8 {
		for i := 0; i < 8 - length; i++ {
			str = "0" + str;
		}	
	}
	return
}