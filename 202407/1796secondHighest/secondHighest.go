package main

func main() {

}

func secondHighest(s string) int {
	ans := make([]int, 10)

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			idx := int(ch) - int('0')
			ans[idx]++
		}
	}
	fidMax := false
	for i := 9; i >= 0; i-- {
		if ans[i] > 0 {
			if !fidMax {
				fidMax = true
				continue
			} else {
				return i
			}
		}
	}
	return -1
}
