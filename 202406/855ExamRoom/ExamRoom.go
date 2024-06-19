package main

func main() {

}

type ExamRoom struct {
	Set []int
	N   int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		Set: make([]int, n),
		N:   n,
	}

}

func (this *ExamRoom) Seat() int {

}

func (this *ExamRoom) Leave(p int) {

}
