package main

func main() {

}

type DataStream struct {
	value int
	k     int
	cnt   int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value: value,
		k:     k,
		cnt:   0,
	}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.value {
		this.cnt++
	} else {
		this.cnt = 0
	}
	return this.cnt >= this.k
}
