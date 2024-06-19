package main

func main() {

}

func scoreOfParentheses(s string) int {
	stark := make([]int, 0)
	for _, ch := range s {
		if byte(ch) == '(' {
			stark = append(stark, 0)
		} else {
			last := stark[len(stark)-1]
			stark = stark[:len(stark)-1]
			add := max(1, last*2)
			if len(stark) == 0 {
				stark = append(stark, add)
			} else {
				pop := stark[len(stark)-1]
				stark = stark[:len(stark)-1]
				stark = append(stark, pop+add)
			}
		}
	}
	return stark[len(stark)-1]
}
