package main

import (
	"math"
)

func main() {

}

func categorizeBox(length int, width int, height int, mass int) string {
	p4 := int(math.Pow10(4))

	b := length >= p4 || width >= p4 || height >= p4 || (length*width*height) >= int(math.Pow10(9))
	h := mass >= 100
	if b && h {
		return "Both"
	}
	if !b && !h {
		return "Neither"
	}
	if b && !h {
		return "Bulky"
	}
	if !b && h {
		return "Heavy"
	}
	return ""
}
