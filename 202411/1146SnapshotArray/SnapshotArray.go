package main

import (
	"sort"
)

func main() {

}

type SnapshotArray struct {
	Data   [][]Node
	SnapId int
}

func Constructor(length int) SnapshotArray {
	s := SnapshotArray{
		Data:   make([][]Node, length),
		SnapId: 0,
	}
	return s
}

func (this *SnapshotArray) Set(index int, val int) {
	ss := this.Data[index]
	no := Node{value: val, snapId: this.SnapId}
	ss = append(ss, no)
	this.Data[index] = ss
}

func (this *SnapshotArray) Snap() int {
	ans := this.SnapId
	this.SnapId++
	return ans
}

// 手写二分的做法
func (this *SnapshotArray) Get1(index int, snapId int) int {
	ss := this.Data[index]
	n := len(ss)
	le, ri := 0, len(ss)
	for le < ri {
		// 找不于等 snapId 的右端点
		mid := le + (ri-le+1)>>1
		if mid >= 0 && mid < n && ss[mid].snapId <= snapId {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	// 检查
	if le < 0 || le >= n || ss[le].snapId > snapId {
		return 0
	}
	return ss[le].value
}

// 调用标准库的做法
func (this *SnapshotArray) Get(index int, snapId int) int {
	ss := this.Data[index]
	n := len(ss)
	le := sort.Search(n, func(i int) bool { return ss[i].snapId > snapId }) - 1
	// 检查
	if le < 0 || le >= n || ss[le].snapId > snapId {
		return 0
	}
	return ss[le].value
}

type Node struct {
	value  int
	snapId int
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
