package main

func main() {

}
func squareIsWhite(coordinates string) bool {
	col := coordinates[0] - '1'
	row := coordinates[1] - 'a'
	if col%2 == 0 {
		return row%2 == 1
	}
	return row%2 == 0
}
