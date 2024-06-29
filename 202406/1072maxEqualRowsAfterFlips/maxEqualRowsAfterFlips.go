package main

func main() {

}

// 如果某两行可以通过这个规则达成一致，要么它们全等，要么它们相反
func maxEqualRowsAfterFlips(matrix [][]int) int {
	cnt := make(map[string]int)
	for _, item := range matrix {
		cnt[gen(item)]++
	}
	ans := 0
	for _, v := range cnt {
		ans = max(ans, v)
	}
	return ans
}

// 统一以1开头，如果是0开头就统一取反
// 那为啥不直接使用取反运算符呢：~ 因为取反会把所有的位都取反，1取反之后不是0，而是 1111111111……0
func gen(item []int) string {
	if item[0] == 0 {
		for i := 0; i < len(item); i++ {
			// 因为只是1和0可以这样做:
			item[i] = item[i] ^ 1
		}
	}
	ans := make([]byte, len(item))
	for i, ch := range item {
		ans[i] = byte(ch + 'a')
	}

	return string(ans)
}
