package main

func main() {

}

func divideString(s string, k int, fill byte) []string {
	start, n := 0, len(s)
	ss := []byte(s)
	ans := make([]string, 0, k)
	for start < n {
		b := min(n, start+k)
		ret := ss[start:b]
		if len(ret) < k {
			i := len(ret)
			for ; i < k; i++ {
				ret = append(ret, fill)
			}
			ans = append(ans, string(ret))
			return ans
		}
		ans = append(ans, string(ret))
		start += k
	}
	return ans
}
