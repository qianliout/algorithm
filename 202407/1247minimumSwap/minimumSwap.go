package main

func main() {

}

func minimumSwap(s1 string, s2 string) int {
	cnt1 := make(map[byte]int)
	for i, ch := range s1 {
		if s1[i] != s2[i] {
			cnt1[byte(ch)]++
		}
	}
	d := cnt1['x'] + cnt1['y']
	if d&1 != 0 {
		return -1
	}
	ans := d / 2
	if cnt1['x']&1 == 1 {
		ans++
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
