package redpacket

import (
	"errors"
	"math/rand/v2"
	"sort"
)

// DistributeCentsProportional 采用“权重和为 1 再乘总金额”的思路：
// 1) 先保证每人至少 1 分（基线），剩余金额按随机权重比例分配；
// 2) 使用“先取地板、再按小数部分大小分配余数”的方式保证总和精确到分；
// 该方法不强制“最大 < 2×最小”比率，可能出现不满足该约束的结果。
func DistributeCentsProportional(totalCents int, n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}
	if totalCents < n {
		return nil, errors.New("totalCents must be >= n (at least 1 cent per person)")
	}

	// 基线：每人先给 1 分
	shares := make([]int, n)
	for i := 0; i < n; i++ {
		shares[i] = 1
	}
	remain := totalCents - n
	if remain == 0 {
		return shares, nil
	}

	// 生成 n 个正随机权重，并归一化
	weights := make([]float64, n)
	sumW := 0.0
	for i := 0; i < n; i++ {
		// 使用 (0,1) 均匀分布的正数作为权重
		w := rand.Float64()
		if w == 0 {
			w = 1e-9 // 极小正数，避免全零导致除零
		}
		weights[i] = w
		sumW += w
	}
	// 按比例给出“分”的初步分配（取地板）并记录小数部分
	type frac struct {
		idx  int
		frac float64
	}
	fracs := make([]frac, n)
	given := 0
	for i := 0; i < n; i++ {
		exact := weights[i] / sumW * float64(remain)
		add := int(exact) // 地板
		shares[i] += add
		given += add
		fracs[i] = frac{idx: i, frac: exact - float64(add)}
	}

	// 剩余的若干分，按照小数部分从大到小依次分配 1 分
	left := remain - given
	sort.Slice(fracs, func(i, j int) bool { return fracs[i].frac > fracs[j].frac })
	for k := 0; k < left; k++ {
		shares[fracs[k%len(fracs)].idx]++
	}

	return shares, nil
}

// DistributeCentsProportionalWithRatio 在比例法基础上增加“最大 < 2×最小”的约束，
// 通过重试若干次（随机权重不同）来尽量满足该约束。若超过重试次数仍无法满足，则返回错误。
func DistributeCentsProportionalWithRatio(totalCents int, n int, maxRetries int) ([]int, error) {
	if maxRetries <= 0 {
		maxRetries = 32
	}
	for attempt := 0; attempt < maxRetries; attempt++ {
		shares, err := DistributeCentsProportional(totalCents, n)
		if err != nil {
			return nil, err
		}
		// 检查比例约束
		minV, maxV := shares[0], shares[0]
		for _, v := range shares {
			if v < minV {
				minV = v
			}
			if v > maxV {
				maxV = v
			}
		}
		if maxV < 2*minV {
			return shares, nil
		}
	}
	return nil, errors.New("failed to satisfy max < 2*min after retries")
}
