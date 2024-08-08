package main

func main() {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})

}

type ATM struct {
	Price []int // 20 ，50 ，100 ，200 和 500 美元
	Data  []int
	N     int
}

func Constructor() ATM {
	return ATM{
		Data:  make([]int, 5),
		N:     5,
		Price: []int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	// 存钱
	for i := 0; i < this.N; i++ {
		this.Data[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	ac := append([]int{}, this.Data...)
	i := this.N - 1
	ans := make([]int, this.N)
	for i >= 0 {
		if amount == 0 {
			break
		}
		a := min(ac[i], amount/this.Price[i])
		// for ac[i] > 0 && this.Price[i] <= amount {
		ac[i] -= a
		amount -= this.Price[i] * a
		ans[i] += a
		// }
		i--
	}
	if amount == 0 {
		this.Data = ac
		return ans
	}
	return []int{-1}
}

// 题目要求没有这么复杂
func (this *ATM) Withdraw2(amount int) []int {
	ac := append([]int{}, this.Data...)
	var dfs func(c int, path []int) bool
	find := make([]int, 0)
	dfs = func(c int, path []int) bool {
		if c < 0 {
			return false
		}
		if c == 0 {
			find = append(find, path...)
			return true
		}
		for i := 0; i < this.N; i++ {
			if this.Data[i] == 0 || this.Price[i] > c {
				continue
			}
			path[i]++
			c -= this.Price[i]
			ac[i]--

			if dfs(c, path) {
				return true
			}
			path[i]--
			c += this.Price[i]
			ac[i]++
		}
		return false
	}
	path := make([]int, 5)
	if !dfs(amount, path) {
		return []int{-1}
	}
	return find
}
