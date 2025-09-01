package redpacket

import (
	"errors"
	"math/rand/v2"
)

// DistributeCents 将总金额（单位：分） totalCents 随机分给 n 个人，返回每个人的金额（单位：分）。
// 约束：
// 1) 总和等于 totalCents
// 2) 每份至少 1 分
// 3) 最大份额 < 2 × 最小份额
// 思路：先按均值向下取整为 base，得到初始均分；然后在不违反“最大 < 2×最小”的前提下，随机分配余额。
// 在输入合法的前提下，算法可以确定终止。
func DistributeCents(totalCents int, n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}
	if totalCents < n { // 至少需要保证每人 1 分
		return nil, errors.New("totalCents must be >= n (at least 1 cent per person)")
	}

	// 从地板除得到基础份额 base 与余额 rem
	base := totalCents / n
	rem := totalCents % n

	// 若 base==0 则与上面的校验矛盾，这里可保证 base>=1
	shares := make([]int, n)
	for i := 0; i < n; i++ {
		shares[i] = base
	}

	// 在“最大 < 2×最小”的约束下随机分配余额。
	// 当前最小值为 base，因此每份最多增长到 limit = 2*base - 1。
	limit := 2*base - 1

	// 若 base==1，则 limit==1，无法再追加任意 1 分；因此必须 rem==0 才能满足约束。
	if base == 1 {
		if rem != 0 {
			return nil, errors.New("cannot satisfy max < 2*min with base=1 and non-zero remainder")
		}
		return shares, nil
	}

	// 收集尚未到达上限的下标
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = i
	}

	for rem > 0 {
		if len(candidates) == 0 {
			// 已无位置可加而不破坏约束
			return nil, errors.New("cannot distribute remainder without breaking ratio constraint")
		}
		idx := candidates[rand.IntN(len(candidates))]
		shares[idx]++
		rem--
		if shares[idx] >= limit { // 达到上限
			// 与末尾交换并截断，从候选集移除
			last := len(candidates) - 1
			for k := 0; k <= last; k++ {
				if candidates[k] == idx {
					candidates[k] = candidates[last]
					candidates = candidates[:last]
					break
				}
			}
		}
	}

	return shares, nil
}

// DistributeYuan 提供以“元”为单位的便捷接口：内部以“分”为单位进行精确分配，再换算为两位小数的元返回。
func DistributeYuan(m float64, n int) ([]float64, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}
	// 四舍五入换算为分
	totalCents := int(m*100 + 0.5)
	cents, err := DistributeCents(totalCents, n)
	if err != nil {
		return nil, err
	}
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = float64(cents[i]) / 100.0
	}
	return res, nil
}
