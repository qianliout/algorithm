package main

func main() {

}

func subtractProductAndSum(n int) int {
	ans := []int{}
	for n > 0 {
		ans = append(ans, n%10)
		n = n / 10
	}
	a, b := 0, 1
	for _, ch := range ans {
		a += ch
		b *= ch
	}
	return b - a
}
