package main

func main() {

}

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}

	a := num / 3
	return []int64{a - 1, a, a + 1}
}
