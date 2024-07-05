package main

func main() {

}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	x, y := coordinates[0][0], coordinates[0][1]
	a, b := coordinates[1][0], coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		ch := coordinates[i]
		c, d := ch[0], ch[1]
		if (a-x)*(d-y) != (c-x)*(b-y) {
			return false
		}
	}

	return true
}
