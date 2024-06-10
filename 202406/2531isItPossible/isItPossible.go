package main

func main() {

}

// 如果可以通过 恰好一次 移动，使 word1 和 word2 中不同字符的数目相等，

func isItPossible(word1 string, word2 string) bool {
	w1, w2 := make(map[int]int), make(map[int]int)
	for i := range word1 {
		w1[int(word1[i]-'a')]++
	}
	for i := range word2 {
		w2[int(word2[i]-'a')]++
	}
	for x, c := range w1 {
		for y, d := range w2 {
			if x == y {
				// 如果相同，其他的字符串相同
				if len(w1) == len(w2) {
					return true
				}
				continue
			}
			// 如果不相同
			// 如果 c==1，那么那 x 移动到 w2之后，那 w1就会少一个
			if len(w1)-b2i(c == 1)+b2i(w1[y] == 0) == len(w2)-b2i(d == 1)+b2i(w2[x] == 0) {
				return true
			}
		}
	}
	return false
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
