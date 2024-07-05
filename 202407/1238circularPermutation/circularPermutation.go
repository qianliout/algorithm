package main

func main() {

}

func circularPermutation(n int, start int) []int {
	res := make([]int, 0)
	for i := 0; i < 1<<n; i++ {
		res = append(res, i^(i>>1))
	}
	j := 0
	for ; j < len(res); j++ {
		if res[j] == start {
			break

		}
	}
	ans := append(res[j:], res[:j]...)
	return ans
}
