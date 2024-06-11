package main

import (
	"sort"
)

func main() {

}

type pair struct{ Value, SnapId int }
type SnapshotArray struct {
	curSnapId int
	history   [][]pair
}

func Constructor(length int) SnapshotArray {
	data := make([][]pair, length)
	return SnapshotArray{history: data, curSnapId: 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.history[index] = append(this.history[index], pair{val, this.curSnapId})
}

func (this *SnapshotArray) Snap() int {
	this.curSnapId++
	return this.curSnapId - 1
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	da := this.history[index]
	// 找快照编号 <= snapId 的最后一次修改记录
	// 等价于找快照编号 >= snapId+1 的第一个修改记录，它的上一个就是答案
	j := sort.Search(len(da), func(k int) bool { return da[k].SnapId >= snapId+1 })
	// 没有找到就返回默认值
	if j == 0 {
		return 0
	}
	return da[j-1].Value
}
