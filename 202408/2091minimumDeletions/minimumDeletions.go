package main

func main() {

}

func minimumDeletions(nums []int) int {
	mx, mi := 0, 0
	for i, ch := range nums {
		if ch > nums[mx] {
			mx = i
		}
		if ch < nums[mi] {
			mi = i
		}
	}
	n := len(nums)
	l := min(mx, mi)
	r := max(mx, mi)
	return min(l+1+n-r, r+1, n-l)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
