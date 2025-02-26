package main

func main() {

}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	random := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if random[node] != nil {
			return random[node]
		}

		ans := &Node{Val: node.Val}
		random[node] = ans
		ans.Next = dfs(node.Next)
		ans.Random = dfs(node.Random)
		return ans
	}

	ans := dfs(head)
	return ans
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
