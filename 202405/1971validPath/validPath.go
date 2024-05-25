package main

func main() {

}

func validPath(n int, edges [][]int, source int, destination int) bool {
	pic := make([][]int, n)
	for _, ch := range edges {
		pic[ch[0]] = append(pic[ch[0]], ch[1])
		pic[ch[1]] = append(pic[ch[1]], ch[0])
	}
	vis := make([]bool, n)
	return dfs(pic, n, vis, source, destination)
}

func dfs(pic [][]int, n int, vis []bool, start, source int) bool {
	if start == source {
		return true
	}
	if vis[start] {
		return false
	}
	ans := false
	vis[start] = true

	nex := pic[start]
	for _, ch := range nex {

		if ch == start {
			continue
		}
		if vis[ch] {
			continue
		}

		ans = ans || dfs(pic, n, vis, ch, source)
	}
	return ans
}
