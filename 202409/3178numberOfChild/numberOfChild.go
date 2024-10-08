package main

func main() {

}

func numberOfChild(n int, k int) int {
	a := k / (n - 1)
	b := k % (n - 1)
	if a&1 == 0 {
		return b
	}
	return n - b - 1
}
