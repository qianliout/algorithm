package main

func main() {

}

func minFlips(target string) int {
	cur := '0'
	cnt := 0
	for _, ch := range target {
		if ch != cur {
			cur = ch
			cnt++
		}
	}
	return cnt
}
