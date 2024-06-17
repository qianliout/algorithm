package main

func main() {

}

type CustomStack struct {
	stark []int
	size  int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stark: make([]int, 0),
		size:  maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stark) >= this.size {
		return
	}
	this.stark = append(this.stark, x)
}

func (this *CustomStack) Pop() int {
	if len(this.stark) == 0 {
		return -1
	}
	n := len(this.stark)
	x := this.stark[n-1]
	this.stark = this.stark[:n-1]
	return x
}

func (this *CustomStack) Increment(k int, va int) {
	n := min(len(this.stark), k)
	for i := 0; i < n; i++ {
		this.stark[i] += va
	}
}
