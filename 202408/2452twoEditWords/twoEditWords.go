package main

func main() {

}

func twoEditWords(queries []string, dictionary []string) []string {
	ans := make([]string, 0)
	for _, q := range queries {
		for _, d := range dictionary {
			if check(q, d) {
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}

// 数据量小，直接暴力
func check(s string, d string) bool {
	// 长度相同
	cnt := 0
	for i := range s {
		if s[i] != d[i] {
			cnt++
		}
	}
	return cnt <= 2
}
