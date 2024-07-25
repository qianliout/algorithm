package main

func main() {

}

func sumGame(num string) bool {
	cnt1, cnt2, sum1, sum2 := 0, 0, 0, 0
	n := len(num)
	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			cnt1++
		} else {
			sum1 += int(num[i]) - int('0')
		}
	}
	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			cnt2++
		} else {
			sum2 += int(num[i]) - int('0')
		}
	}

	// 如果是奇数，最后一手是 alice，他总是能弄的不相等
	if (cnt1+cnt2)%2 != 0 {
		return true
	}
	// 最后一手是 bob 那么 alice 只能把差值弄的特别大（大于9）这个 bob 才能弄相等

	sub := 9*(cnt1-cnt2)/2 + (sum1 - sum2)
	return sub != 0
}
