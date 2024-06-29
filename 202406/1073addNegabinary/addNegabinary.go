package main

func main() {

}

// 最简单的方法，转成10进制，相加之后再转 -2进制
// 但是有数位超限制,所以不能过，
func addNegabinary(arr1 []int, arr2 []int) []int {
	return baseNeg2(to10(arr1) + to10(arr2))
}

func to10(arr1 []int) int {
	ans := 0
	for i := 0; i < len(arr1); i++ {
		ans = ans*-2 + arr1[i]
	}
	return ans
}

func baseNeg2(n int) []int {
	ans := make([]int, 0)
	if n == 0 {
		return []int{0}
	}
	for n != 0 {
		mod := n % -2
		n = n / -2
		if mod == -1 {
			n++
			mod = 1
		}
		ans = append(ans, mod)
	}

	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}
	return ans
}
