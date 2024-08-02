package main

func main() {

}

type Robot struct {
	w   int
	h   int
	s   int
	mod int
}

func Constructor(width int, height int) Robot {
	return Robot{
		w:   width,
		h:   height,
		mod: (width + height - 2) * 2,
		s:   0,
	}
}

func (this *Robot) Step(num int) {
	this.s = (this.s+num-1)%this.mod + 1
}

func (this *Robot) GetPos() []int {
	a, b, _ := this.get()
	return []int{a, b}
}

func (this *Robot) GetDir() string {
	_, _, d := this.get()
	return d
}

func (this *Robot) get() (int, int, string) {
	w, h, s := this.w, this.h, this.s
	if s <= w-1 {
		return s, 0, "East"
	}
	if s <= w+h-2 {
		return w - 1, s - w + 1, "North"
	}
	if s <= w*2+h-3 {
		return w*2 + h - 3 - s, h - 1, "West"
	}
	return 0, (w+h-2)*2 - s, "South"
}
