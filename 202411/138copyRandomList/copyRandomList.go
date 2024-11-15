package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	cnt := make(map[*Node]*Node)

	res := &Node{}

	cur := res
	first := head
	for first != nil {
		node := &Node{Val: first.Val}
		cur.Next = node
		cnt[first] = node
		first = first.Next
		cur = cur.Next
	}
	cur = res.Next
	first = head

	for first != nil {
		if first.Random != nil {
			cur.Random = cnt[first.Random]
		}
		cur = cur.Next
		first = first.Next
	}
	return res.Next
}

// 怎么理解呢
func copyRandomList(head *Node) *Node {
	mem := make(map[*Node]*Node)
	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := mem[node]; ok {
			return v
		}
		no := &Node{Val: node.Val}
		mem[node] = no
		no.Next = dfs(node.Next)
		no.Random = dfs(node.Random)
		return no
	}
	ans := dfs(head)
	return ans
}
