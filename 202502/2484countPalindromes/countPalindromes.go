package main

func main() {

}

func countPalindromes(s string) int {
	ss := []byte(s)
	rev := revers([]byte(s))
	suf1 := make([]int, 10)
	suf2 := make([]int, 100)
	for _, ch := range rev {
		a := int(ch) - int('0')
		// 假设 a= 1
		// 这一步是算01，11，21，31，41，51，61，71，81，91 有多少个
		for j, c := range suf1 {
			suf2[j*10+a] += c
		}
		// 这里要注意，一定是先算上面的
		suf1[a]++
	}
	ans := 0
	pre1, pre2 := make([]int, 10), make([]int, 100)
	for _, ch := range ss {
		a := int(ch) - int('0')
		// 因为要把 a 当做中间的数，所以要把 a 的所有统计数据撤销
		suf1[a]--
		for j, c := range suf1 {
			suf2[j*10+a] -= c
		}
		for i, c1 := range pre2 {
			ans += c1 * suf2[i]
		}

		for j, c := range pre1 {
			pre2[j*10+a] += c
		}
		// 这里要注意，一定是先算上面的
		pre1[a]++

	}
	return ans % 1000000007
}

func revers(ss []byte) []byte {
	l, r := 0, len(ss)-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}
	return ss
}
