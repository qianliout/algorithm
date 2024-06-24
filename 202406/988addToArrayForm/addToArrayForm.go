package main

func main() {

}

func addToArrayForm(num []int, k int) []int {
	numK := gen(k)
	return add(num, numK)
}

func add(nums1, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return add(nums2, nums1)
	}
	ad := 0

	n := len(nums1)
	sub := len(nums1) - len(nums2)
	ans := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i-sub >= 0 {
			a := (ad + nums1[i] + nums2[i-sub]) % 10
			ad = (ad + nums1[i] + nums2[i-sub]) / 10
			ans[i+1] = a
		} else {
			a := (ad + nums1[i]) % 10
			ad = (ad + nums1[i]) / 10
			ans[i+1] = a
		}
	}
	if ad > 0 {
		ans[0] = ad
	}
	if ans[0] == 0 {
		return ans[1:]
	}
	return ans
}

func gen(num int) []int {
	ans := make([]int, 0)
	for num > 0 {
		ans = append(ans, num%10)
		num = num / 10
	}
	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}

	return ans
}
