package main

func main() {

}

func getSmallestString(s string, k int) string {
	ss := []byte(s)
	for i, ch := range ss {
		mi := min(int(ch)-int('a'), int('z')-int(ch)+1)
		if mi > k {
			ss[i] = byte(int(ch) - k)
			break
		} else {
			ss[i] = 'a'
			k -= mi
		}
	}

	return string(ss)
}
