package main

func main() {

}

type BrowserHistory struct {
	Page []string
	Pos  int // 当前页的位置
}

func Constructor(homepage string) BrowserHistory {
	c := BrowserHistory{
		Page: []string{homepage},
		Pos:  0,
	}
	return c
}

/*
使用一个栈记录浏览历史，使用一个 pos 记录当前网页在栈中的位置。每次 back 和 forward 操作都只更新 pos 。
因为visit操作会把浏览历史前进的记录全部删除，所以每次 visit 先根据 pos 更新下栈顶指针，然后再将 url 入栈。
*/
func (this *BrowserHistory) Visit(url string) {
	this.Page = this.Page[:this.Pos+1]
	this.Page = append(this.Page, url)
	this.Pos++
}

func (this *BrowserHistory) Back(steps int) string {
	idx := max(0, this.Pos-steps)
	this.Pos = idx
	return this.Page[idx]
}

func (this *BrowserHistory) Forward(steps int) string {
	idx := min(len(this.Page)-1, this.Pos+steps)
	this.Pos = idx
	return this.Page[idx]
}
