package main

func main() {

}

func isWinner(player1 []int, player2 []int) int {
	n1 := cal(player1)
	n2 := cal(player2)
	if n1 == n2 {
		return 0
	}
	if n1 > n2 {
		return 1
	}
	return 2
}

func cal(nums []int) int {
	ans := 0
	for i, ch := range nums {
		if i >= 1 && nums[i-1] == 10 {
			ans += 2 * ch
			continue
		}
		if i >= 2 && nums[i-2] == 10 {
			ans += 2 * ch
			continue
		}
		ans += ch
	}
	return ans
}
