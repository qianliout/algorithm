package main

func main() {

}

type Allocator struct {
	Data []int
	N    int
}

func Constructor(n int) Allocator {
	s := Allocator{
		Data: make([]int, n),
		N:    n,
	}
	return s
}

func (this *Allocator) Allocate(size int, mID int) int {
	for i := 0; i < this.N; i++ {
		if this.Check(i, size) {
			for j := i; j < i+size; j++ {
				this.Data[j] = mID
			}
			return i
		}
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	cnt := 0
	for i := 0; i < this.N; i++ {
		if this.Data[i] == mID {
			cnt++
			this.Data[i] = 0
		}
	}
	return cnt
}

func (this *Allocator) Check(start, size int) bool {
	if start+size > this.N {
		return false
	}

	for i := start; i < start+size; i++ {
		if this.Data[i] != 0 {
			return false
		}
	}
	return true
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
