package main

func main() {

}

type FreqStack struct {
	Stark [][]int
	Freq  map[int]int
}

func Constructor() FreqStack {
	return FreqStack{
		Stark: make([][]int, 0),
		Freq:  make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	c := this.Freq[val]
	if c == len(this.Stark) {
		this.Stark = append(this.Stark, []int{val})

	} else {
		this.Stark[c] = append(this.Stark[c], val)
	}
	this.Freq[val]++
}

func (this *FreqStack) Pop() int {
	for {
		for len(this.Stark[len(this.Stark)-1]) == 0 {
			this.Stark = this.Stark[:len(this.Stark)-1]
		}
		if len(this.Stark) == 0 {
			break
		}
		n := len(this.Stark)
		m := len(this.Stark[n-1])
		pop := this.Stark[n-1][m-1]
		this.Stark[n-1] = this.Stark[n-1][:m-1]
		this.Freq[pop]--
		return pop
	}
	return 0
}
