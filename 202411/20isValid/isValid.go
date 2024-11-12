package main

func main() {

}

// 这个方法不对 ："([)]"
func isValid1(s string) bool {
	a, b, c := 0, 0, 0
	for _, ch := range s {
		switch byte(ch) {
		case '(':
			a++
		case '{':
			b++
		case '[':
			c++
		case ')':
			a--
		case '}':
			b--
		case ']':
			c--
		}

		if a < 0 || b < 0 || c < 0 {
			return false
		}
	}
	return a == 0 && b == 0 && c == 0
}
func isValid(s string) bool {
	cnt := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	st := make([]byte, 0)
	for _, ch := range s {
		if ch == '[' || ch == '(' || ch == '{' {
			st = append(st, byte(ch))
			continue
		}
		if len(st) == 0 {
			return false
		}
		last := st[len(st)-1]
		if cnt[byte(ch)] != last {
			return false
		}
		st = st[:len(st)-1]
	}

	return len(st) == 0
}
