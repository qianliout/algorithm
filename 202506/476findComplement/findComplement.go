package main

import "math/bits"

func findComplement(num int) int {
	// #分三步
	// 1、num.bit_length() 先计算二进制位长度
	// 2、(1<<num.bit_length()) -1 得到当前长度二进制位全为1的表示,1000 -1 = 0111
	// 3、与原数异或即为答案, 111 ^ 101 = 010
	n := bits.Len(uint(num))
	return num ^ (1<<n - 1)
}
