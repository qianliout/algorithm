package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 17))
}

func fullJustify(words []string, maxWidth int) []string {
	res := make([][]string, 0)
	i, n := 0, len(words)
	for i < n {
		lev := make([]string, 0)
		cnt := 0
		i0 := i
		for i0 < n {
			cnt += len(words[i0])
			if cnt > maxWidth {
				break
			}
			lev = append(lev, words[i0])
			cnt += 1
			i0++
		}
		res = append(res, lev)
		i = i0
	}
	ans := make([]string, 0)
	for k := 0; k < len(res)-1; k++ {
		ans = append(ans, help(res[k], maxWidth))
	}
	ans = append(ans, last(res[len(res)-1], maxWidth))
	return ans
}

func help(strs []string, mx int) string {
	cnt := 0
	for _, ch := range strs {
		cnt += len(ch)
	}
	n := len(strs)
	if n == 1 {
		return strs[0] + strings.Repeat(" ", mx-cnt)
	}

	j, k := (mx-cnt)%(n-1), (mx-cnt)/(n-1)
	ans := make([]string, n-1)
	for i := range ans {
		if i < j {
			ans[i] = strings.Repeat(" ", k+1)
		} else {
			ans[i] = strings.Repeat(" ", k)
		}
	}
	res := make([]string, 0)
	for i := 0; i < n-1; i++ {
		res = append(res, strs[i])
		res = append(res, ans[i])
	}
	res = append(res, strs[n-1])
	s := strings.Join(res, "")
	return s
}

func last(strs []string, mx int) string {
	s := strings.Join(strs, " ")
	s = s + strings.Repeat(" ", mx-len(s))
	return s
}
