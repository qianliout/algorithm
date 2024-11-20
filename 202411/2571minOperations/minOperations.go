package main

func main() {

}

func minOperations(n int) int {

	var dfs func(i int) int
	mem := make(map[int]int)
	dfs = func(x int) int {
		if x&(x-1) == 0 {
			return 1
		}
		if v, ok := mem[x]; ok {
			return v
		}
		// 最后一个位是1的数
		lb := x & -x
		// 此时我们可以加上这个数，或者减去这个数
		ans := 1 + min(dfs(x+lb), dfs(x-lb))
		mem[x] = ans
		return ans
	}
	ans := dfs(n)
	return ans
}

// 2的幂 说明这个数的二进制位只有一个1,这说明我们可以每次只考虑最后的一位1
