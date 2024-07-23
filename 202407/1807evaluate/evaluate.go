package main

import (
	"strings"
)

func main() {

}

// s 中不会有嵌套括号对。
func evaluate(s string, knowledge [][]string) string {
	kn := make(map[string]string)
	for _, ch := range knowledge {
		k, v := ch[0], ch[1]
		kn[k] = v
	}

	ss := []byte(s)
	start := 0
	ans := make([]string, 0)

	for i, ch := range ss {
		if ch == '(' {
			ans = append(ans, string(ss[start:i]))
			start = i + 1
		}
		if ch == ')' {
			key := string(ss[start:i])
			v := kn[key]
			if v == "" {
				v = "?"
			}
			ans = append(ans, v)
			start = i + 1
		}
	}

	ans = append(ans, string(ss[start:]))

	return strings.Join(ans, "")
}
