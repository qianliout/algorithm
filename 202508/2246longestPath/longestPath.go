package main

func main() {

}

func longestPath2(parent []int, s string) int {
	n := len(parent)

	// 第一步：构建邻接表，将parent数组转换为树的结构
	// g[i] 存储节点i的所有子节点
	g := make([][]int, n)
	for i, p := range parent {
		if p == -1 { // 根节点，跳过
			continue
		}
		// 将节点i添加为节点p的子节点
		g[p] = append(g[p], i)
	}

	// ans用于记录全局最长路径的长度（边数）
	ans := 0

	// DFS函数定义：返回以节点i为根的子树中，从i出发向下的最长有效路径长度（边数）
	// 有效路径：相邻节点字符不相同
	var dfs func(i int) int

	dfs = func(i int) int {
		// res: 从当前节点i向下的最长有效路径长度
		res := 0

		// 遍历当前节点i的所有子节点
		for _, ch := range g[i] {
			// 递归计算子节点ch为根的最长路径长度，+1是因为要加上从i到ch的边
			child := dfs(ch) + 1

			// 只有当子节点字符与当前节点字符不同时，才能形成有效路径
			if s[ch] != s[i] {
				/*
					关键理解：这里在计算"经过当前节点i的最长路径"
					路径可能的形状：子树1 -> i -> 子树2

					res: 当前已知的从i向下的最长路径长度
					child: 从i到ch子树的最长路径长度
					res + child: 就是经过i连接两个不同子树的路径长度

					为什么顺序很重要：
					1. 先用旧的res计算经过i的路径长度：ans = max(ans, res + child)
					2. 再更新res为更长的单边路径：res = max(res, child)

					这样确保我们考虑了所有可能经过节点i的路径组合
				*/
				ans = max(ans, res+child) // 更新全局最长路径（可能经过当前节点i）
				res = max(res, child)     // 更新从i向下的最长单边路径
			}
			// 如果s[ch] == s[i]，则跳过这个子节点，因为不能形成有效路径
		}

		// 返回从当前节点i向下的最长有效路径长度
		return res
	}

	// 从根节点0开始DFS
	dfs(0)

	// 返回路径上的节点数量，所以要+1（边数+1=节点数）
	return ans + 1
}

/*
问题描述：
给你一棵 树（即一个连通、无向、无环图），根节点是节点 0 ，这棵树由编号从 0 到 n - 1 的 n 个节点组成。用下标从 0 开始、长度为 n 的数组 parent 来表示这棵树，其中 parent[i] 是节点 i 的父节点，由于节点 0 是根节点，所以 parent[0] == -1 。
另给你一个字符串 s ，长度也是 n ，其中 s[i] 表示分配给节点 i 的字符。
请你找出路径上任意一对相邻节点都没有分配到相同字符的 最长路径 ，并返回该路径的长度。

算法核心思想举例：
假设有树：
    0(a)
   /    \
  1(b)  2(c)
 /  \    \
3(c) 4(b) 5(a)

parent = [-1, 0, 0, 1, 1, 2], s = "abccba"

执行过程：
1. 构建邻接表：g[0]=[1,2], g[1]=[3,4], g[2]=[5]

2. DFS(0):
   - DFS(1):
     - DFS(3): 返回0（叶节点）
     - DFS(4): 返回0（叶节点）
     - 对于child=3: s[3]='c' != s[1]='b', child=0+1=1
       ans = max(0, 0+1) = 1, res = max(0, 1) = 1
     - 对于child=4: s[4]='b' == s[1]='b', 跳过
     - 返回res=1

   - DFS(2):
     - DFS(5): 返回0（叶节点）
     - 对于child=5: s[5]='a' != s[2]='c', child=0+1=1
       ans = max(1, 0+1) = 1, res = max(0, 1) = 1
     - 返回res=1

   - 在DFS(0)中：
     - 对于child=1: s[1]='b' != s[0]='a', child=1+1=2
       ans = max(1, 0+2) = 2, res = max(0, 2) = 2
     - 对于child=2: s[2]='c' != s[0]='a', child=1+1=2
       ans = max(2, 2+2) = 4, res = max(2, 2) = 2
     - 返回res=2

3. 最终ans=4，返回4+1=5（节点数）

最长路径是：3(c) -> 1(b) -> 0(a) -> 2(c) -> 5(a)，长度为5个节点
*/

func longestPath(parent []int, s string) int {
	n := len(parent)

	// 第一步：构建邻接表，将parent数组转换为树的结构
	// g[i] 存储节点i的所有子节点
	g := make([][]int, n)
	for i, p := range parent {
		if p == -1 { // 根节点，跳过
			continue
		}
		// 将节点i添加为节点p的子节点
		g[p] = append(g[p], i)
	}

	// ans用于记录全局最长路径的长度（边数）
	ans := 0

	// DFS函数定义：返回以节点i为根的子树中，从i出发向下的最长有效路径长度（边数）
	// 有效路径：相邻节点字符不相同
	var dfs func(i int) int

	dfs = func(i int) int {
		// res: 从当前节点i向下的最长有效路径长度
		res := 1

		// 遍历当前节点i的所有子节点
		for _, ch := range g[i] {
			// 递归计算子节点ch为根的最长路径长度，+1是因为要加上从i到ch的边
			child := dfs(ch)

			// 只有当子节点字符与当前节点字符不同时，才能形成有效路径
			if s[ch] != s[i] {
				/*
					关键理解：这里在计算"经过当前节点i的最长路径"
					路径可能的形状：子树1 -> i -> 子树2

					res: 当前已知的从i向下的最长路径长度
					child: 从i到ch子树的最长路径长度
					res + child: 就是经过i连接两个不同子树的路径长度

					为什么顺序很重要：
					1. 先用旧的res计算经过i的路径长度：ans = max(ans, res + child)
					2. 再更新res为更长的单边路径：res = max(res, child)

					这样确保我们考虑了所有可能经过节点i的路径组合
				*/
				ans = max(ans, res+child) // 更新全局最长路径（可能经过当前节点i）
				res = max(res, child)     // 更新从i向下的最长单边路径
			}
			// 如果s[ch] == s[i]，则跳过这个子节点，因为不能形成有效路径
		}

		// 返回从当前节点i向下的最长有效路径长度
		return res
	}

	// 从根节点0开始DFS
	dfs(0)

	// 返回路径上的节点数量，所以要+1（边数+1=节点数）
	return ans + 1
}
