package main

func main() {

}

func balancedStringSplit(s string) int {
	ans := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ans++
		}
	}
	return ans
}
