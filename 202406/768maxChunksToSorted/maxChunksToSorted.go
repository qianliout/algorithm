package main

func main() {

}

func maxChunksToSorted(arr []int) int {
	stark := make([]int, 0)
	for _, ch := range arr {
		if len(stark) > 0 && stark[len(stark)-1] > ch {
			head := stark[len(stark)-1]
			for len(stark) > 0 && stark[len(stark)-1] > ch {
				stark = stark[:len(stark)-1]
			}
			stark = append(stark, head)
		} else {
			stark = append(stark, ch)
		}
	}
	return len(stark)
}
