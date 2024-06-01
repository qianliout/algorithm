package main

func main() {

}

// 如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。
func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				ans++
			}
		}
	}
	return ans
}
