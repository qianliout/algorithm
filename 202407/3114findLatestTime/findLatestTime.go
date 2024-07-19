package main

func main() {

}

func findLatestTime(s string) string {
	ss := []byte(s)
	for i, ch := range ss {
		if byte(ch) != '?' {
			continue
		}
		switch i {
		case 0:
			if ss[1] != '?' && ss[1] >= '2' {
				ss[i] = '0'
			} else {
				ss[i] = '1'
			}
		case 1:
			if ss[0] == '0' {
				ss[1] = '9'
			} else {
				ss[1] = '1'
			}
		case 3:
			ss[i] = '5'
		case 4:
			ss[i] = '9'
		}
	}
	return string(ss)
}
