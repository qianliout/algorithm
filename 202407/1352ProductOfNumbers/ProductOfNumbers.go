package main

import (
	"fmt"
)

func main() {
	c := Constructor()
	c.Add(3)
	c.Add(0)
	c.Add(2)
	c.Add(5)
	c.Add(4)
	fmt.Println(c.GetProduct(2))
	fmt.Println(c.GetProduct(3))
	fmt.Println(c.GetProduct(4))
}

type ProductOfNumbers struct {
	Product []int
	Data    []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{Product: []int{1}, Data: make([]int, 0)}

}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.Product = []int{1}
		return
	}
	n := len(this.Product)
	this.Product = append(this.Product, this.Product[n-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {

	if k >= len(this.Product) { // 说明中途遇到0了
		return 0
	}
	ri := len(this.Product) - 1
	le := ri - k

	return this.Product[ri] / this.Product[le]
}
