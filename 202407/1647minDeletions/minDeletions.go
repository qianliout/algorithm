package main

func main() {

}

func minDeletions(s string) int {
	cnt := make(map[byte]int)
	for _, ch := range s {
		cnt[byte(ch)]++
	}
	freq := make(map[int]int)
	for _, ch := range cnt {
		freq[ch]++
	}
	ans := 0

	for k, v := range freq {
		if v == 1 {
			continue
		}
		del := 0
		for ; k <= v; k++ {
			if freq[k] == 0 {
				break
			}
		}

	}

}
