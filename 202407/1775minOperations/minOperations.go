package main

func main() {

}

func minOperations(nums1 []int, nums2 []int) int {
	if 6*len(nums1) < len(nums2) || 6*len(nums2) < len(nums1) {
		return -1
	}
	d := sum(nums2) - sum(nums1)
	if d == 0 {
		return d
	}
	if d < 0 {
		return minOperations(nums2, nums1)
		// nums1, nums2 = nums2, nums1
		// d = -d
	}
	ans := 0
	// 数据范围 1 <= nums1[i], nums2[i] <= 6
	cnt := make([]int, 6)
	for _, ch := range nums1 {
		cnt[6-ch]++
	}
	for _, ch := range nums2 {
		cnt[ch-1]++
	}
	// 贪心的做法
	for i := 5; i >= 0; i-- {
		if i*cnt[i] >= d {
			// 向上取整
			ans += (d + i - 1) / i
			break
		}
		ans += cnt[i]
		d -= i * cnt[i]
	}

	return ans
}

func sum(nums []int) int {
	ans := 0
	for _, ch := range nums {
		ans += ch
	}
	return ans
}
