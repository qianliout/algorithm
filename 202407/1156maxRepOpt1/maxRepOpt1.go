package main

func main() {

}

func maxRepOpt1(text string) int {
	cnt := make(map[byte]int)
	for _, ch := range text {
		cnt[byte(ch)]++
	}
	res, i, j, n := 0, 0, 0, len(text)
	for i < n {
		for text[j] == text[i] {
			j++
		}
		res = max(res, j-i)
		// 遇到不同的了，就想着交换一次
		k := j + 1
		for text[k] == text[i] {
			k++
		}
		// 这里为啥要取一个min呢，因为题目中说了可以交换一次，但是得用和 i相同的交换，如果前面已经i用完了，就不能交换了
		res = max(res, min(k-i, cnt[byte(text[i])]))
		i = j
	}
	return res
}
