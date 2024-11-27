package main

import (
	"sort"
)

func main() {

}

type TimeMap struct {
	Data map[string][]*Node
	Key  map[string]map[string]*Node
}

func Constructor() TimeMap {
	data := make(map[string][]*Node)
	key := make(map[string]map[string]*Node)
	return TimeMap{
		Data: data,
		Key:  key,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	var no *Node
	if this.Key[key] != nil && this.Key[key][value] != nil {
		no = this.Key[key][value]
	}
	if no != nil {
		no.tm = timestamp
	} else {
		no = &Node{key: key, value: value, tm: timestamp}
		if this.Key[key] == nil {
			this.Key[key] = make(map[string]*Node)
		}
		this.Key[key][value] = no
		this.Data[key] = append(this.Data[key], no)
	}
	// 排序
	ss := this.Data[key]
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].tm < ss[j].tm
	})
	this.Data[key] = ss
}

func (this *TimeMap) Get(key string, timestamp int) string {
	ss := this.Data[key]
	n := len(ss)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= 0 && mid < n && ss[mid].tm <= timestamp {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	if le < 0 || le >= n || ss[le].tm > timestamp {
		return ""
	}
	return ss[le].value

}

type Node struct {
	key   string
	value string
	tm    int
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
