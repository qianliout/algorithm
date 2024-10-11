package main

func main() {

}

func findPrefixScore(nums []int) []int64 {
	mx, s := 0, 0
	ans := make([]int64, 0)
	for _, ch := range nums {
		mx = max(mx, ch)
		s += mx + ch
		ans = append(ans, int64(s))
	}
	return ans
}
