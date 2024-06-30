package main

func main() {

}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	t := []byte("!@#$%^&*()-+")
	tm := make(map[byte]bool)
	for _, ch := range t {
		tm[ch] = true
	}

	a, b, c, d := false, false, false, false
	for i, ch := range password {
		if ch >= 'a' && ch <= 'z' {
			a = true
		}
		if ch >= 'A' && ch <= 'Z' {
			b = true
		}
		if ch >= '0' && ch <= '9' {
			c = true
		}
		if tm[byte(ch)] {
			d = true
		}
		if i > 0 && password[i] == password[i-1] {
			return false
		}
	}
	return a && b && c && d
}
