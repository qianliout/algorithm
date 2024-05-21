package main

func main() {

}

func xorQueries(arr []int, queries [][]int) []int {
	sum := make([]int, len(arr)+1)
	for i, ch := range arr {
		sum[i+1] = sum[i] ^ ch
	}
	ans := make([]int, len(queries))
	for i, ch := range queries {
		le, ri := ch[0], ch[1]
		ans[i] = sum[ri+1] ^ sum[le]
	}
	return ans
}
