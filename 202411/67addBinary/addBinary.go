package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary1(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	m, n := len(a), len(b)
	aa, bb := []byte(a), []byte(b)
	i, j := 0, len(aa)-1
	for i < j {
		aa[i], aa[j] = aa[j], aa[i]
		i++
		j--
	}
	i, j = 0, len(bb)-1
	for i < j {
		bb[i], bb[j] = bb[j], bb[i]
		i++
		j--
	}
	add := 0
	ans := make([]byte, 0)
	for k := 0; k < m; k++ {
		sum := int(aa[k]-'0') + add
		add = 0
		if k < n {
			sum += int(bb[k] - '0')
		}
		if sum >= 2 {
			add = 1
			sum -= 2
		}
		ans = append(ans, byte('0'+sum))
	}
	if add > 0 {
		ans = append(ans, byte('0'+add))
	}
	i, j = 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	m, n := len(a), len(b)
	ans := make([]byte, m+1)
	add := 0
	for k := m - 1; k >= 0; k-- {
		num := int(a[k]-'0') + add
		add = 0
		if k-(m-n) >= 0 {
			num += int(b[k-(m-n)] - '0')
		}
		if num >= 2 {
			add = 1
			num -= 2
		}
		ans[k+1] = byte('0' + num)
	}
	if add > 0 {
		ans[0] = byte('0' + add)
		return string(ans)
	}

	return string(ans[1:])
}
