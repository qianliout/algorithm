package main

func main() {

}

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num = num >> 1
	}
	return res
}
