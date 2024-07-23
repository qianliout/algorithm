package main

func main() {

}

func findTheWinner(n int, k int) int {
	x := 0
	for i := 1; i <= n; i++ {
		x = (x + k) % i
	}
	return x + 1
}
