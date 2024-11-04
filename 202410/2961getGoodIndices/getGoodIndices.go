package main

func main() {

}

func getGoodIndices(variables [][]int, target int) []int {
	n := len(variables)
	pairs := make([]Pair, n)
	for i, ch := range variables {
		pairs[i] = Pair{
			A: ch[0],
			B: ch[1],
			C: ch[2],
			M: ch[3],
		}
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if cal(pairs[i]) == target {
			ans = append(ans, i)
		}
	}
	return ans
}

type Pair struct {
	A, B, C, M int
}

func cal(p Pair) int {
	a := pow(p.A, p.B, 10)
	b := pow(a, p.C, p.M)
	return b
}

func pow(a, b, m int) int {
	if b == 0 {
		return 1
	}
	c := pow(a, b/2, m)
	if b%2 == 1 {
		return (a * c * c) % m
	}
	return (c * c) % m
}
