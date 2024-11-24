package main

func main() {

}

func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n = n / 5
	}
	return ans
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	i := 1
	for ; i*i <= x; i++ {
	}
	return i
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	b := myPow(x, n/2)
	if n%2 == 0 {
		return b * b
	}
	return b * b * x
}

func maxPoints(points [][]int) int {
	n := len(points)
	ans := 1
	for i, ch := range points {
		x1, y1 := ch[0], ch[1]
		for j := i + 1; j < n; j++ {
			ch2 := points[j]
			x2, y2 := ch2[0], ch2[1]
			cnt := 2
			for k := j + 1; k < n; k++ {
				ch3 := points[k]
				x3, y3 := ch3[0], ch3[1]
				// 这个判断是技巧
				if (y2-y1)*(x3-x2) == (y3-y2)*(x2-x1) {
					cnt++
				}
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}

type pair struct {
	x, y int
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
