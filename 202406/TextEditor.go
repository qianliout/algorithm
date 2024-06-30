package main

import (
	"fmt"
)

func main() {
	te := Constructor()
	te.AddText("leetcode")
	fmt.Println(te.DeleteText(4))
	te.AddText("practice")
	fmt.Println(te.CursorRight(3))
	fmt.Println(te.CursorLeft(8))
	fmt.Println(te.DeleteText(10))
	fmt.Println(te.CursorLeft(2))
	fmt.Println(te.CursorRight(6))
	// [null,null,4,null,"etpractice","leet",4,"","practi"]...
}

type TextEditor struct {
	left  []byte
	right []byte
}

func Constructor() TextEditor {
	return TextEditor{
		left:  make([]byte, 0),
		right: make([]byte, 0),
	}
}

func (this *TextEditor) AddText(text string) {
	this.left = append(this.left, []byte(text)...)
}

func (this *TextEditor) DeleteText(k int) int {
	k0 := k
	for len(this.left) > 0 && k > 0 {
		this.left = this.left[:len(this.left)-1]
		k--
	}
	return k0 - k
}

func (this *TextEditor) CursorLeft(k int) string {
	for len(this.left) > 0 && k > 0 {
		this.right = append(this.right, this.left[len(this.left)-1])
		this.left = this.left[:len(this.left)-1]
		k--
	}
	return this.text()
}

func (this *TextEditor) CursorRight(k int) string {
	for len(this.right) > 0 && k > 0 {
		this.left = append(this.left, this.right[len(this.right)-1])
		this.right = this.right[:len(this.right)-1]
		k--
	}
	return this.text()
}

func (this *TextEditor) text() string {
	mx := max(0, len(this.left)-10)
	return string(this.left[mx:])
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
