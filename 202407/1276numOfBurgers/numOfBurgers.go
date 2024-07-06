package main

func main() {

}

func numOfBurgers(t int, c int) []int {

	ans := make([]int, 0)
	if t < 2*c || t > 4*c {
		return ans
	}
	if (t-2*c)%2 != 0 {
		return ans
	}
	a := (t - 2*c) / 2
	b := (t - a*4) / 2
	return []int{a, b}
}
