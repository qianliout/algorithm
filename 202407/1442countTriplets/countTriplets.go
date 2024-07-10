package main

func main() {

}

func countTriplets(arr []int) int {
	n := len(arr)
	pre := make([]int, n+1)
	for i, ch := range arr {
		pre[i+1] = pre[i] ^ ch
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			for k := j; k <= n; k++ {
				if pre[j-1]^pre[i-1] == pre[k]^pre[j-1] {
					ans++
				}
			}
		}
	}
	return ans
}
