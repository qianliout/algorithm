package main

func main() {

}

// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	hash := make(map[int]int)
	ans := make([]int, len(nums1))
	st := make([]int, 0)
	for i := range ans {
		ans[i] = -1
	}
	for i, c := range nums2 {
		for len(st) > 0 && c > nums2[st[len(st)-1]] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			hash[nums2[last]] = c
		}
		st = append(st, i)
	}
	for i, c := range nums1 {
		if v, ok := hash[c]; ok {
			ans[i] = v
		}
	}
	return ans
}
