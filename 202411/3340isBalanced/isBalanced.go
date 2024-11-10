package main

func main() {

}

func isBalanced(num string) bool {
	a, b, n := 0, 0, len(num)
	for i := 0; i < n; i = i + 2 {
		a += int(num[i]) - int('0')
		if i+1 < n {
			b += int(num[i+1]) - int('0')
		}
	}
	return a == b
}
