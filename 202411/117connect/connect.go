package main

func main() {

}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		lev := make([]*Node, 0)
		for _, no := range queue {
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		for i := 0; i < len(lev)-1; i++ {
			lev[i].Next = lev[i+1]
		}
		queue = lev
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
