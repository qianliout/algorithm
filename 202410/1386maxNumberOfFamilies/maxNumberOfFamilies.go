package main

func main() {

}
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	cnt := map[int]int{}
	for _, v := range reservedSeats {
		r, c := v[0], v[1]
		if c == 1 || c == 10 {
			continue
		}
		cnt[r] += 1 << (9 - c)
	}
	ans := n * 2
	for _, v := range cnt {
		// v & 0xF0 == 0：检查 v 的高 4 位是否全为 0。
		// v & 0x3C == 0：检查 v 的第 2 到第 5 位（从右往左数）是否全为 0。
		// v & 0x0F == 0：检查 v 的低 4 位是否全为 0。
		// 0xF0 是一个十六进制数，对应的十进制数值是 240。在二进制表示中，0xF0 是 11110000。
		// 这个值常用于位操作，特别是在处理字节或位掩码时。具体到你的代码中，
		// v & 0xF0 用于提取 v 的高 4 位

		// 0x3C 是一个十六进制数，对应的十进制数值是 60。在二进制表示中，0x3C 是 00111100
		// 0x0F 是一个十六进制数，对应的十进制数值是 15。在二进制表示中，0x0F 是 00001111
		if v&0xF0 == 0 || v&0x3C == 0 || v&0x0F == 0 {
			ans--
		} else {
			ans -= 2
		}
	}

	return ans
}
