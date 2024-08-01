package main

func main() {

}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	cnt := make(map[int]int)
	nums1, nums2, nums3 = dup(nums1), dup(nums2), dup(nums3)
	for _, ch := range nums1 {
		cnt[ch]++
	}
	for _, ch := range nums2 {
		cnt[ch]++
	}
	for _, ch := range nums3 {
		cnt[ch]++
	}
	ans := make([]int, 0)
	for k, v := range cnt {
		if v >= 2 {
			ans = append(ans, k)
		}
	}
	return ans
}

func dup(nums []int) []int {
	cnt := make(map[int]int)
	for _, ch := range nums {
		cnt[ch]++
	}
	ans := make([]int, 0)
	for k := range cnt {
		ans = append(ans, k)
	}
	return ans
}
