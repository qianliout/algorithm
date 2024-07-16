package main

func main() {

}

type FrontMiddleBackQueue struct {
	Data1 []int // 前 前面最多比后面多1个元shu
	Data2 []int // 后

}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		Data1: make([]int, 0),
		Data2: make([]int, 0),
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {

}

func (this *FrontMiddleBackQueue) PushBack(val int) {

}

func (this *FrontMiddleBackQueue) PopFront() int {

}

func (this *FrontMiddleBackQueue) PopMiddle() int {

}

func (this *FrontMiddleBackQueue) PopBack() int {

}
