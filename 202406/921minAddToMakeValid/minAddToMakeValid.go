package main

func main() {

}

func minAddToMakeValid(s string) int {
	stark := make([]int, 0)
	ans := 0
	for i := range s {
		if s[i] == '(' {
			stark = append(stark, i)
		} else {
			if len(stark) == 0 {
				ans++
				continue
			} else {
				stark = stark[:len(stark)-1]
			}
		}

	}
	return ans + len(stark)
}
