package main

func flipAndInvertImage(image [][]int) [][]int {
	for _, ch := range image {
		l, r := 0, len(ch)-1
		for l < r {
			ch[l], ch[r] = ch[r], ch[l]
			l++
			r--
		}
	}
	for _, ch := range image {
		for i, j := range ch {
			ch[i] = j ^ 1
		}
	}
	return image
}
