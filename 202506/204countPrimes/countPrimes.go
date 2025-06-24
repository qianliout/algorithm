package main

func main() {

}

// 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
func countPrimes(n int) int {
	// return Count1(n)
	return Count2(n)
}

func Count1(n int) int {
	prim := make([]bool, n)
	for i := 2; i < n; i++ {
		prim[i] = true
	}
	for i := 2; i < n; i++ {
		if prim[i] == true {
			for j := 2; j*i < n; j++ {
				prim[i*j] = false
			}
		}
	}
	ans := 0
	for _, j := range prim {
		if j {
			ans++
		}
	}
	return ans
}

func Count2(n int) int {
	notPrim := make([]bool, n)
	prim := make([]int, 0)
	for i := 2; i < n; i++ {
		if !notPrim[i] {
			prim = append(prim, i)
		}
		for _, j := range prim {
			if i*j >= n {
				break
			}
			notPrim[i*j] = true
			if i%j == 0 {
				break
			}

		}
	}
	return len(prim)
}
