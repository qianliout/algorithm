package main

func main() {

}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m, n := len(nums1), len(nums2)
	ans := make([][]int, 0)
	i, j := 0, 0
	for {
		if i == m {
			ans = append(ans, nums2[j:]...)
			return ans
		}
		if j == n {
			ans = append(ans, nums1[i:]...)
			return ans
		}
		if nums1[i][0] == nums2[j][0] {
			ans = append(ans, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] > nums2[j][0] {
			ans = append(ans, nums2[j])
			j++
		} else if nums1[i][0] < nums2[j][0] {
			ans = append(ans, nums1[i])
			i++
		}
	}
}
