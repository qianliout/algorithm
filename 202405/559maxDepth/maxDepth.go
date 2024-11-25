package main

func main() {

}

/**
 * Definition for a Head.
 * type Head struct {
 *     Val int
 *     Children []*Head
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		ans++
		lev := make([]*Node, 0)
		for _, no := range queue {
			for _, ch := range no.Children {
				if ch != nil {
					lev = append(lev, ch)
				}
			}
		}
		queue = lev
	}
	return ans
}
