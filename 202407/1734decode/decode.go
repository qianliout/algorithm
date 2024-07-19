package main

func main() {

}

func decode(encoded []int) []int {
	n := len(encoded) + 1
	a := 0
	for i := 0; i < len(encoded); i += 2 {
		a = a ^ encoded[i]
	}
	for i := 1; i <= n; i++ {
		a = a ^ i
	}
	ans := make([]int, n)
	ans[n-1] = a
	for i := len(ans) - 2; i >= 0; i-- {
		ans[i] = ans[i+1] ^ encoded[i]
	}
	return ans
}

// 给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个 奇数 。 全排列是重点
