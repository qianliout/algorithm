package main

func main() {

}

func findPermutationDifference(s string, t string) int {
	ss1, ss2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s); i++ {
		ss1[int(s[i]-'a')] = i
		ss2[int(t[i]-'a')] = i
	}
	ans := 0
	for i := 0; i < 26; i++ {
		ans += abs(ss1[i] - ss2[i])
	}
	return ans
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
