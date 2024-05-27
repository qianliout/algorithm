package main

func main() {

}

/*
给你一个 有向无环图 ， n 个节点编号为 0 到 n-1 ，以及一个边数组 edges ，其中 edges[i] = [fromi, toi] 表示一条从点  fromi 到点 toi 的有向边。
找到最小的点集使得从这些点出发能到达图中所有点。题目保证解存在且唯一。
你可以以任意顺序返回这些节点编号。
*/
// 统计入度为0的点集
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	in := make([][]int, n)
	for i := range in {
		in[i] = make([]int, 0)
	}
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		in[y] = append(in[y], x)
	}
	ans := make([]int, 0)
	for k, ch := range in {
		if len(ch) == 0 {
			ans = append(ans, k)
		}
	}
	return ans
}
