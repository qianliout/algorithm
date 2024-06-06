package main

import "fmt"

func main() {
	fmt.Println(uniqueLetterString("ABC"))
	fmt.Println(uniqueLetterString("ABA"))
	fmt.Println(uniqueLetterString("LEETCODE"))
}

/*
将所有子串按照其末尾字符的下标分组。
*/
func uniqueLetterString(s string) int {
	first := make(map[byte]int) // 上上次出现的位置
	sec := make(map[byte]int)   // 上一次出现的位置
	// 初值,这里这个技巧很巧妙
	for i := 'A'; i <= 'Z'; i++ {
		first[byte(i)] = -1
		sec[byte(i)] = -1
	}
	total := 0
	ans := 0
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		c := ss[i]
		// 假设s =  BCADEAFGA c 是最后的 A
		// 假设把 c 加到 以ss[i-1]结尾的子串上，会在 ss[i-1]的基础上 增加或减少（用这个理解 total:因为是在前一个子串的基础上的结果）
		// 1,c 单独作为一个子串，会增加1
		// 2,如果加在 G，和 FG 后面，会让 G 和 FG 这两个子串的数增加一
		// 3,如果加在 AFG,EAFG,DEAFG的后面，会减少1
		// 4,再向前走就没有影响
		total += 1 + (i - 1 - sec[c]) - (sec[c] - first[c])
		// total += i - 2*sec[c] + first[c]

		ans += total
		first[c] = sec[c]
		sec[c] = i
	}
	return ans
}
