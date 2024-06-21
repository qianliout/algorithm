package main

import (
	"strings"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	path := make([]string, 0)
	var dfs func(root *TreeNode, target int) bool
	dfs = func(root *TreeNode, target int) bool {
		if root == nil {
			return false
		}
		if root.Val == target {
			return true
		}

		path = append(path, "L")
		if dfs(root.Left, target) {
			return true
		}
		// path = path[:len(path)-1]
		// path = append(path, "R")
		// 上面两步合成下面的一步
		path[len(path)-1] = "R"
		if dfs(root.Right, target) {
			return true
		}
		path = path[:len(path)-1]
		return false
	}
	dfs(root, startValue)
	leftPath := append([]string{}, path...)
	path = path[:0]
	dfs(root, destValue)
	rightPath := append([]string{}, path...)

	for len(leftPath) > 0 && len(rightPath) > 0 && leftPath[0] == rightPath[0] {
		leftPath = leftPath[1:]
		rightPath = rightPath[1:]
	}
	return strings.Repeat("U", len(leftPath)) + strings.Join(rightPath, "")
}
