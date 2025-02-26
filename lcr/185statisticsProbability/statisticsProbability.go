package main

import (
	"math"
)

func main() {

}

func statisticsProbability2(num int) []float64 {
	f := make([]float64, 0)  // 一个骰子有6面
	f1 := make([]float64, 6) // 只有一个骰子的情况
	for i := 0; i < 6; i++ {
		f1[i] = float64(1) / float64(6)
	}
	f = append(f, f1...)
	for i := 2; i <= num; i++ {
		// 确定这里数组的大小是个难点
		tem := make([]float64, 5*i+1)
		for j := 0; j < len(f); j++ {
			for k := 0; k < 6; k++ {
				tem[j+k] += f[j] / float64(6)
			}
		}
		f = tem
	}
	return f
}

func statisticsProbability(num int) []float64 {
	f := make([]float64, 0)  // 一个骰子有6面
	f1 := make([]float64, 6) // 只有一个骰子的情况
	for i := 0; i < 6; i++ {
		f1[i] = float64(1) / float64(6)
	}
	// 这样写会超过内存
	inf := int(math.Pow(6, float64(num)))
	f = append(f, f1...)
	for i := 2; i <= num; i++ {
		tem := make([]float64, inf)
		for j := 0; j < len(f); j++ {
			for k := 0; k < 6; k++ {
				tem[j+k] += f[j] / float64(6)
			}
		}
		// 把为0的去掉
		for k := len(tem) - 1; k >= 0; k-- {
			if tem[k] != 0 {
				tem = tem[:k+1]
				break
			}
		}

		f = tem
	}
	return f
}

// 每个骰子摇到 1 至 6 的概率相等，都为
// 将每个骰子的点数看作独立情况，共有 6 种「点数组合」。例如 n=2 时的点数组合为：
// (1,1),(1,2),⋯,(2,1),(2,2),⋯,(6,1),⋯,(6,6)
// n 个骰子「点数和」的范围为 [n,6n] ，数量为 6n−n+1=5n+1 种。
