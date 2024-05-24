package main

import (
	"math"
)

func main() {

}

type pair struct {
	day   int
	price int
}

type StockSpanner struct {
	stark  []pair
	curDay int
}

func Constructor() StockSpanner {
	stark := make([]pair, 0)
	stark = append(stark, pair{
		day:   -1,
		price: math.MaxInt32,
	})
	// 写入一个初值，防止stark空
	return StockSpanner{stark: stark, curDay: -1}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.stark) > 0 && this.stark[len(this.stark)-1].price <= price {
		this.stark = this.stark[:len(this.stark)-1]
	}
	this.curDay++
	this.stark = append(this.stark, pair{price: price, day: this.curDay})
	return this.curDay - this.stark[len(this.stark)-2].day
}
