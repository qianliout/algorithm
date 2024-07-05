package main

func main() {

}

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(f func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	le, ri := 1, z
	// [1,4],[4,1]都算答案，所以这里要控制边界
	for ri >= 1 && le <= z {
		a := f(le, ri)
		if a == z {
			ans = append(ans, []int{le, ri})
			le++
			ri--
		} else if a < z {
			le++
		} else {
			ri--
		}
	}
	return ans
}
