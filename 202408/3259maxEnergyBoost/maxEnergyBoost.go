package main

func main() {

}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	eng := [][]int{energyDrinkA, energyDrinkB}
	n := len(energyDrinkA)
	var dfs func(i, j int) int
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 2)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		if mem[i][j] >= 0 {
			return mem[i][j]
		}
		ans := max(dfs(i-1, j), dfs(i-2, j^1)) + eng[j][i]
		mem[i][j] = ans
		return ans
	}
	return int64(max(dfs(n-1, 0), dfs(n-1, 1)))
}
