package main

func main() {

}

func numJewelsInStones(jewels string, stones string) int {
	exist := make(map[byte]bool)
	for _, ch := range jewels {
		exist[byte(ch)] = true
	}
	cnt := 0
	for _, ch := range stones {
		if exist[byte(ch)] {
			cnt++
		}
	}
	return cnt
}
