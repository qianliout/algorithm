package main

func main() {

}

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	a := int(coordinate1[0]) + int(coordinate1[1])
	b := int(coordinate2[0]) + int(coordinate2[1])
	return a&1 == b&1
}
