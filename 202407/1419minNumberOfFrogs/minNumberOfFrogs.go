package main

func main() {

}

func minNumberOfFrogs(croakOfFrogs string) int {
	var previous = map[byte]byte{
		'c': 'k',
		'r': 'c',
		'o': 'r',
		'a': 'o',
		'k': 'a',
	}
	cnt := make(map[byte]int)
	for _, ch := range croakOfFrogs {
		pre := previous[byte(ch)]
		// 如果有前一个，说明还没有叫完
		// 题目要求 返回模拟字符串中所有蛙鸣所需不同青蛙的最少数目。croak
		// 所以当一个青蛙叫完之后，可以继续叫，在代码中的表达就是：当一个叫到c时，如明前面的 k 大0也就是有一个青蛙刚叫完，那就让这个青蛙接着叫
		if cnt[pre] > 0 {
			cnt[pre]--
		} else if ch != 'c' { // cnt[pre]==0 && ch!='c'
			return -1
		}
		cnt[byte(ch)]++
	}
	// 所有的青蛙都必须叫完
	if cnt['c'] > 0 || cnt['r'] > 0 || cnt['o'] > 0 || cnt['a'] > 0 {
		return -1
	}
	// 最后只能发出 k
	return cnt['k']
}
