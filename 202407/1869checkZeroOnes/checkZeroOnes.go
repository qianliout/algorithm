package main

func main() {

}

func checkZeroOnes(s string) bool {
	return cal(s, '1') > cal(s, '0')
}

func cal(s string, b byte) int {
	ans := 0
	i := 0
	for i < len(s) {
		for i < len(s) && byte(s[i]) != b {
			i++
		}
		j := i
		for ; j < len(s); j++ {
			if byte(s[j]) != b {
				break
			}
		}
		ans = max(ans, j-i)
		i = j
	}
	return ans
}
