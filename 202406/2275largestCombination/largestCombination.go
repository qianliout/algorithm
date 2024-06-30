package main

func main() {

}

func largestCombination(candidates []int) int {
	ans := 0
	for i := 0; i < 64; i++ {
		cnt := 0
		for _, ch := range candidates {
			if (ch>>i)&1 == 1 {
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return ans
}
