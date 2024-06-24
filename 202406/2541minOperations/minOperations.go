package main

func main() {

}

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	if len(nums1) != len(nums2) {
		return -1
	}
	if k == 0 {
		if same(nums1, nums2) {
			return 0
		}
		return -1
	}

	n := len(nums1)
	sub := make([]int, n)
	for i := 0; i < n; i++ {
		if (nums2[i]-nums1[i])%k != 0 {
			return -1
		}
		sub[i] = (nums2[i] - nums1[i]) / k
	}
	sum := 0
	ans := 0
	for _, ch := range sub {
		sum += ch
		if ch > 0 {
			ans += ch
		}
	}
	if sum != 0 {
		return -1
	}

	return int64(ans)
}
func same(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
