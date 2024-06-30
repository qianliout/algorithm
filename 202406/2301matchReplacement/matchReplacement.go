package main

func main() {

}

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	mapp := make(map[byte][]byte)
	for _, ch := range mappings {
		mapp[ch[0]] = append(mapp[ch[0]], ch[1])
	}
	cur := []byte(sub)
	n := len(s)
	le, ri := 0, 0
	wind := make([]byte, 0)
	for le <= ri && ri < n {
		wind = append(wind, s[ri])

		for len(wind) > len(sub) {
			wind = wind[1:]
		}
		if check(wind, cur, mapp) {
			return true
		}
		ri++
	}
	return false
}

func check(pre, cur []byte, mapp map[byte][]byte) bool {
	if len(pre) != len(cur) {
		return false
	}
	for i := 0; i < len(cur); i++ {
		if cur[i] == pre[i] {
			continue
		}
		flag := false
		for _, ch := range mapp[cur[i]] {
			if pre[i] == ch {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}

	return true
}
