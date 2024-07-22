package main

func main() {

}

func minOperations(s string) int {
	start := byte('0')
	cnt1 := 0
	for _, ch := range s {
		if byte(ch) == start {
			cnt1++
		}
		start = change(start)
	}

	start = byte('1')
	cnt2 := 0
	for _, ch := range s {
		if byte(ch) == start {
			cnt2++
		}
		start = change(start)
	}

	return min(cnt1, cnt2)
}

func change(start byte) byte {
	if start == '0' {
		return '1'
	}
	return '0'
}
