package main

func main() {

}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 横
	if rec2[0] >= rec1[2] || rec2[2] <= rec1[0] {
		return false
	}
	if rec2[1] >= rec1[3] || rec2[3] <= rec1[1] {
		return false
	}
	return true
}
