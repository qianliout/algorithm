package main

func main() {

}

func countKConstraintSubstrings(s string, k int) int {
	ans := 0
	cnt := make([]int, 2)
	left, right, n := 0, 0, len(s)
	for right < n {
		ch := int(s[right]) % 2
		cnt[ch]++
		right++
		for left <= right && cnt[0] > k && cnt[1] > k {
			le := int(s[left]) % 2
			cnt[le]--
			left++
		}
		ans += right - left
	}
	return ans
}
