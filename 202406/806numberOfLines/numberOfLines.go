package main

func main() {

}

func numberOfLines(widths []int, s string) []int {
	a, b := 0, 0
	for _, t := range s {
		idx := int(t) - int('a')
		if b+int(widths[idx]) > 100 {
			a++
			b = widths[idx]
		} else {
			b += widths[idx]
		}
	}
	if b > 0 {
		a++
	}
	return []int{a, b}
}
