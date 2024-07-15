package main

func main() {

}

type Node struct {
	Name  string
	Child []*Node
}

type ThroneInheritance struct {
	Node      *Node
	Data      map[string]*Node
	DeathName map[string]bool // 名字不相同，所以可以这样写
}

func Constructor(kingName string) ThroneInheritance {
	no := &Node{
		Name:  kingName,
		Child: make([]*Node, 0),
	}
	c := ThroneInheritance{
		Node:      no,
		Data:      make(map[string]*Node),
		DeathName: make(map[string]bool),
	}
	c.Data[no.Name] = no
	return c
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	node := this.Data[parentName]
	no := &Node{Name: childName}
	node.Child = append(node.Child, no)
	this.Data[no.Name] = no
}

func (this *ThroneInheritance) Death(name string) {
	this.DeathName[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	return this.dfs(this.Node)
}

func (this *ThroneInheritance) dfs(root *Node) []string {
	if root == nil {
		return []string{}
	}
	res := make([]string, 0)
	if !this.DeathName[root.Name] {
		res = append(res, root.Name)
	}
	for i := range root.Child {
		ret := this.dfs(root.Child[i])
		res = append(res, ret...)
	}
	return res
}
