package main

import (
	"container/heap"
)

func main() {

}

type FoodRatings struct {
	Rat map[string]int
	Cui map[string]string
	Hm  map[string]MinHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	f := FoodRatings{
		Rat: make(map[string]int),
		Hm:  make(map[string]MinHeap),
		Cui: make(map[string]string),
	}
	n := len(foods)
	for i := 0; i < n; i++ {
		pa := Pair{
			Fo: foods[i],
			Ra: ratings[i],
			CU: cuisines[i],
		}
		f.Rat[pa.Fo] = pa.Ra
		f.Cui[pa.Fo] = pa.CU
		hm := f.Hm[pa.CU]
		if hm == nil {
			hm = make(MinHeap, 0)
		}
		heap.Push(&hm, pa)
		f.Hm[pa.CU] = hm
	}
	return f
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	cu := this.Cui[food]

	this.Rat[food] = newRating

	mh := this.Hm[cu]
	heap.Push(&mh, Pair{Fo: food, Ra: newRating, CU: cu})
	this.Hm[cu] = mh
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	hm := this.Hm[cuisine]
	for hm.Len() > 0 {
		fi := hm[0]
		if fi.Ra != this.Rat[fi.Fo] {
			heap.Pop(&hm)
			continue
		}
		this.Hm[cuisine] = hm
		return fi.Fo
	}
	return ""
}

type Pair struct {
	Fo string
	Ra int
	CU string
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].Ra != h[j].Ra {
		return h[i].Ra > h[j].Ra
	}
	return h[i].Fo < h[j].Fo
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
