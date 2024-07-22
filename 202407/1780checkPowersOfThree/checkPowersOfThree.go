package main

func main() {

}

func checkPowersOfThree(n int) bool {
	for n > 0 {
		// n%3==1 是因为 x 的0次方是1
		if n%3 == 0 || n%3 == 1 {
			n = n / 3
			continue
		}
		return false
	}

	return true
}
