package main

import (
	"math/bits"
)

type Fancy struct {
	sz int
	tr seg
}

const n = 1e5 + 5

func Constructor() Fancy {
	t := make(seg, 2<<bits.Len(uint(n-1)))
	a := make([]data, n)
	t.build(a, 1, 1, n)
	return Fancy{sz: 0, tr: t}
}

func (this *Fancy) Append(val int) {
	this.sz += 1
	this.tr.update(1, this.sz, this.sz, pair{val, 1})
}

func (this *Fancy) AddAll(inc int) {
	if this.sz == 0 {
		return
	}
	this.tr.update(1, 1, this.sz, pair{inc, 1})
}

func (this *Fancy) MultAll(m int) {
	if this.sz == 0 {
		return
	}
	this.tr.update(1, 1, this.sz, pair{0, m})
}

func (this *Fancy) GetIndex(idx int) int {
	if idx+1 > this.sz {
		return -1
	}
	return this.tr.query(1, idx+1, idx+1).v
}

/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */

const mod = 1e9 + 7

var todoInit = pair{0, 1}

type data struct{ v int }
type pair struct{ add, mul int }
type seg []struct {
	l, r int
	d    data
	todo pair
}

// 线段树模板，只需要实现 mergeInfo 和 do，其余都是固定的
func mergeInfo(a, b data) data {
	return data{(a.v + b.v) % mod}
}

func (t seg) do(O int, p pair) {
	o := &t[O]

	sz := o.r - o.l + 1
	// 先乘后加
	if p.mul != 1 {
		o.d.v = o.d.v * p.mul % mod
		o.todo.add = o.todo.add * p.mul % mod
		o.todo.mul = o.todo.mul * p.mul % mod
	}
	if p.add != 0 {
		o.d.v = (o.d.v + sz*p.add) % mod
		o.todo.add = (o.todo.add + p.add) % mod
	}
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != todoInit {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit
	}
}

func (t seg) build(a []data, o, l, r int) {
	t[o].l, t[o].r = l, r

	t[o].todo = todoInit
	if l == r {
		t[o].d = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r int, v pair) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func (t seg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].d
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}
