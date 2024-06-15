package main

func main() {

}

func numSubmat(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				pre[i+1][j+1] = 0
			} else {
				pre[i+1][j+1] = pre[i+1][j] + 1
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			col := pre[i][j]
			// todo(liuqiang) 这里的计算方法还是没有理解透彻
			for k := i; i >= 1; k-- {
				col = min(col, pre[k][j])
				if col == 0 {
					break
				}
				ans += col
			}
		}
	}
	return ans
}
