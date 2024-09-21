package main

func main() {

}

type LUPrefix struct {
	x int
	s map[int]struct{}
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		x: 1,
		s: map[int]struct{}{},
	}
}

func (this *LUPrefix) Upload(video int) {
	this.s[video] = struct{}{}
}

// 上传进度不会减少
func (this *LUPrefix) Longest() int {
	for {
		if _, ok := this.s[this.x]; ok {
			this.x++
		} else {
			break
		}
	}
	return this.x - 1
}
