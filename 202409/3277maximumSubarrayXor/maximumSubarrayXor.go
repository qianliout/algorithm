package main

func main() {

}

func maximumSubarrayXor(nums []int, queries [][]int) []int {
	// 定义 f[i][j] 表示下标从 i 到 j 的子数组的「数组的异或值」，根据上面的讨论，有
	// f[i][j]=f[i][j−1]⊕f[i+1][j]
	// 为了回答询问，我们需要计算下标从 i 到 j 的子数组中的所有子数组的 f 值的最大值，将其记作 mx[i][j]。
	// mx[i][j]=max(f[i][j],mx[i][j−1],mx[i+1][j])
	n := len(nums)
	f := make([][]int, n)
	mx := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		mx[i] = make([]int, n)
	}

	// 初值
	// for i := 0; i < n; i++ {
	// 	f[i][i] = nums[i]
	// 	mx[i][i] = nums[i]
	// }
	// 这里一定要好好理解为啥要倒序,因为i是由i和i+1转移过来的，在计算 i 时得保证 i+1 计算完成
	for i := n - 1; i >= 0; i-- {
		f[i][i] = nums[i]
		mx[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i][j-1] ^ f[i+1][j]
			mx[i][j] = max(f[i][j], mx[i][j-1], mx[i+1][j])
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = mx[q[0]][q[1]]
	}

	return ans
}
