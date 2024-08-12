package main

func main() {

}

func oddString(words []string) string {
	cnt := make(map[sub]int)
	cnt2 := make(map[sub]string)
	for _, ch := range words {
		s := cal(ch)
		cnt[s]++
		cnt2[s] = ch
	}
	for k, v := range cnt {
		if v == 1 {
			return cnt2[k]
		}
	}
	return ""
}

func cal(s string) sub {
	ans := [21]int{}
	for i := 1; i < len(s); i++ {
		ans[i-1] = int(s[i]) - int(s[i-1])
	}
	return ans
}

type sub [21]int
