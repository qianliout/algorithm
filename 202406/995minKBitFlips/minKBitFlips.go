package main

func main() {

}

// 直接暴力会超时
func minKBitFlips(nums []int, k int) int {
	n := len(nums)
	queue := make([]int, 0)
	res := 0
	for i, x := range nums {
		if len(queue) > 0 && queue[0]+k <= i {
			queue = queue[1:]
		}
		if x == 1 {
			if len(queue)%2 == 0 {
				continue
			}
			if i+k > n {
				return -1
			}
			queue = append(queue, i)
			res++
		}
		if x == 0 {
			if len(queue)%2 == 1 {
				continue
			}
			if i+k > n {
				return -1
			}
			queue = append(queue, i)
			res++
		}

		// 这样写更简洁但是不容易理解
		// if len(queue)%2 == x {
		// 	if i+k > n {
		// 		return -1
		// 	}
		// 	queue = append(queue, i)
		// 	res++
		// }
	}
	return res
}
