package main

func main() {

}

func countPoints(rings string) int {
	cnt := make([]int, 10)
	n := len(rings)
	for i := 0; i < n; i = i + 2 {
		c := byte(rings[i])
		idx := int(rings[i+1] - '0')
		switch c {
		case 'R':
			cnt[idx] = cnt[idx] | (1 << 0)
		case 'G':
			cnt[idx] = cnt[idx] | (1 << 1)
		case 'B':
			cnt[idx] = cnt[idx] | (1 << 2)
		}
	}
	ans := 0
	for _, v := range cnt {
		if v == 7 {
			ans++
		}
	}
	return ans
}
