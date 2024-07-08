package main

func main() {

}

type Cashier struct {
	N        int
	Dis      int
	Products map[int]Product
	Num      int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	c := Cashier{
		Num:      0,
		N:        n,
		Dis:      discount,
		Products: make(map[int]Product),
	}

	for i := 0; i < len(prices); i++ {
		c.Products[products[i]] = Product{ID: products[i], Pri: prices[i]}
	}
	return c

}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.Num = this.Num + 1
	var all float64
	for i, pro := range product {
		all += float64(this.Products[pro].Pri * amount[i])
	}

	if this.Num == this.N {
		this.Num = 0
		all = all - float64(this.Dis)*all/100
	}
	return all
}

type Product struct {
	ID  int
	Pri int
}
