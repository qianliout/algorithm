package main

func main() {

}

func makeEqual(words []string) bool {
	n := len(words)
	all := 0
	cnt := make([]int, 26)
	for i := range words {
		for _, ch := range words[i] {
			all++
			idx := int(ch) - int('a')
			cnt[idx]++
		}
	}
	if all%n != 0 {
		return false
	}
	for _, ch := range cnt {
		if ch%n != 0 {
			return false
		}
	}
	return true
}
