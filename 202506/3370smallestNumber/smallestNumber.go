package main

func main() {

}

func smallestNumber(n int) int {
	i := n
	for {
		if check(i) {
			return i
		}
		i++
	}
}

func check(a int) bool {
	for a > 0 {
		if a&1 == 0 {
			return false
		}
		a = a >> 1
	}
	return true
}
