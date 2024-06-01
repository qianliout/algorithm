package main

func main() {

}

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	idx := 0
	for idx < n-1 {
		if bits[idx] == 0 {
			idx++
		} else {
			idx += 2
		}
	}
	return idx == n-1
}
