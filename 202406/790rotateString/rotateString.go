package main

func main() {

}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	s = s + s

	ss := []byte(s)
	le, ri := 0, len(goal)
	for le < ri && ri < len(ss) {
		if string(ss[le:ri]) == goal {
			return true
		}
		le++
		ri++
	}
	return false
}
