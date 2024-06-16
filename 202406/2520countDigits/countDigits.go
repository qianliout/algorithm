package main

func main() {

}

func countDigits(num int) int {
	bits := make([]int, 10)
	a := num
	for a > 0 {
		b := a % 10
		bits[b]++
		a = a / 10
	}
	ans := 0
	for i := 0; i <= 9; i++ {
		if bits[i] == 0 {
			continue
		}
		if num%i == 0 {
			ans += bits[i]
		}
	}
	return ans
}
