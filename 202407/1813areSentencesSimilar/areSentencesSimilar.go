package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	fmt.Println(areSentencesSimilar("of", "A lot of words"))
}

// 可以通过往其中一个句子插入一个任意的句子
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	a := strings.Split(sentence1, " ")
	b := strings.Split(sentence2, " ")
	return f(a, b)

}

func f(a, b []string) bool {
	if len(a) < len(b) {
		a, b = b, a
	}
	n, m := len(a), len(b)
	i, j := 0, 0
	for i < m && a[i] == b[i] {
		i++
	}
	for j < m && a[n-1-j] == b[m-1-j] {
		j++
	}

	return i+j >= m
}
