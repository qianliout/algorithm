package main

func main() {

}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	tyM := make(map[string]int)
	coM := make(map[string]int)
	nameM := make(map[string]int)
	for _, ch := range items {
		ty, co, na := ch[0], ch[1], ch[2]
		tyM[ty]++
		coM[co]++
		nameM[na]++
	}
	switch ruleKey {
	case "type":
		return tyM[ruleValue]
	case "color":
		return coM[ruleValue]
	case "name":
		return nameM[ruleValue]
	}
	return 0
}
