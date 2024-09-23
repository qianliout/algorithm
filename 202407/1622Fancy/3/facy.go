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
		// 在乘法更新时通常不需要使用 sz（区间大小），因为乘法操作是对节点当前存储的值进行的直接操作，而不是累加操作。对于线段树中的一个节点来说，它存储的值通常是其覆盖范围内的某些聚合信息（如总和、最小值、最大值等）。
		// 在这个特定的函数中，o.d.v 表示的是节点的某种聚合值。当执行乘法更新时，我们只需要将这个聚合值乘以给定的因子 p.mul 即可。这是因为乘法操作不会受到区间长度的影响——无论区间内有多少个元素，每个元素都应当被相同的因子所乘。
		// 例如，如果 o.d.v 存储的是一个区间的元素之和，那么当我们需要将区间内每个元素都乘以某个常数 k 时，我们只需将 o.d.v 乘以 k，而不需要知道具体的元素数量。
		// 因此，在乘法更新时不使用 sz 是合理的。而加法更新时使用 sz 是因为我们需要将区间内每个元素都加上一个常数值，所以需要知道区间的实际长度来正确地更新聚合值
		o.d.v = o.d.v * p.mul % mod

		// 在乘法更新时更新加法因子是为了保证后续的加法操作能够正确地反映之前的所有操作。具体原因如下：
		// 懒惰传播：线段树中的懒惰传播机制允许我们延迟一些操作，直到真正需要的时候才执行。这意味着在某些节点上可能会累积多个待执行的操作。
		// 组合操作：假设某个节点已经有一个待执行的加法操作 add 和一个待执行的乘法操作 mul。当新的乘法操作 p.mul 应用到这个节点时，我们需要确保之前的加法操作也能够正确地应用。
		// 具体来说：
		// 当前节点的待执行加法操作 o.todo.add 需要乘以新的乘法因子 p.mul，以确保在最终执行加法操作时，结果仍然正确。
		// 例如：
		// 假设当前节点有 o.todo.add = 5 和 o.todo.mul = 2。
		// 新的乘法操作 p.mul = 3 应用到节点上。
		// 更新后的加法操作应该是 o.todo.add = 5 * 3 = 15，这样在最终执行加法操作时，每个元素先乘以 3 再加 15，而不是先加 5 再乘以 3。
		// 因此，在乘法更新时更新加法因子是为了确保所有操作的顺序和效果正确无误。这样可以避免在后续的加法操作中出现错误的结果。
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
