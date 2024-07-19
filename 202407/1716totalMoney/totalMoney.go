package main

func main() {

}

func totalMoney(n int) int {
	start := 1
	ans := 0
	for n > 0 {
		for i := start; i < start+7 && n > 0; i++ {
			ans += i
			n--
		}
		start++
	}
	return ans
}
