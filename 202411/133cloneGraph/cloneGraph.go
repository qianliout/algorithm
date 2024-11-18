package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var dfs func(o *Node) *Node

	mem := make(map[*Node]*Node)

	dfs = func(o *Node) *Node {
		if o == nil {
			return nil
		}
		if v, ok := mem[o]; ok {
			return v
		}
		nn := &Node{Val: o.Val, Neighbors: make([]*Node, len(o.Neighbors))}
		// 一定要在这里就要记录对应关系
		mem[o] = nn
		for i, ch := range o.Neighbors {
			nn.Neighbors[i] = dfs(ch)
		}
		// 不能到这里才记录
		// mem[o] = nn
		return nn
	}
	nn := dfs(node)
	return nn
}
