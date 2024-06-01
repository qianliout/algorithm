package main

import (
	"fmt"
)

func main() {
	fmt.Println(shoppingOffers([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}))
	fmt.Println(shoppingOffers([]int{2, 2}, [][]int{{1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 1, 4}, {1, 1, 5}, {1, 1, 6}, {1, 1, 7}, {1, 1, 8}, {1, 1, 9}, {1, 1, 10}, {1, 1, 11}, {1, 1, 12}, {1, 1, 13}, {1, 1, 14}, {1, 1, 15}}, []int{10, 10}))
}

/*
输入：price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
输出：14
解释：有 A 和 B 两种物品，价格分别为 ¥2 和 ¥5 。
大礼包 1 ，你可以以 ¥5 的价格购买 3A 和 0B 。
大礼包 2 ，你可以以 ¥10 的价格购买 1A 和 2B 。
需要购买 3 个 A 和 2 个 B ， 所以付 ¥10 购买 1A 和 2B（大礼包 2），以及 ¥4 购买 2A 。
返回 确切 满足购物清单所需花费的最低价格，你可以充分利用大礼包的优惠活动。你不能购买超出购物清单指定数量的物品，
即使那样会降低整体价格。任意大礼包可无限次购买。
*/

func shoppingOffers(price []int, special [][]int, needs []int) int {
	cost := 0
	path := make([]int, len(price))
	// 先计算购买单品的价格，如果这个部分用dfs的话，太耗时
	for k, v := range needs {
		cost += v * price[k]
	}
	dfs(price, special, needs, path, 0, &cost, 0)
	return cost
}

// path:当前购物总数；cur:当前消费金额; cost:最小消费金额；pos：当前礼包位置
func dfs(price []int, special [][]int, needs []int, path []int, cur int, cost *int, pos int) {
	if same(needs, path) {
		if cur < *cost {
			*cost = cur
		}
		return
	}

	for k, spec := range special {
		// 不允许选择前面的礼包，避免出现礼包1 2, 2 1重复计算的情况
		if k < pos {
			continue
		}
		// 这个大礼包更贵，这个就没有用
		if cur+spec[len(spec)-1] >= *cost {
			continue
		}
		flag := true

		for sp := 0; sp < len(spec)-1; sp++ {
			// 这个礼包里有我们不需要的物品
			if sp >= len(path) {
				flag = false
				break
			}
			// 选择这个大礼包之后，数量多了，就不能选
			// 如果礼包里的商品数量大于待购清单，放弃该礼包
			if path[sp]+spec[sp] > needs[sp] {
				flag = false
				break
			}
		}
		if flag {
			// 买的数据数量
			for i := 0; i < len(price); i++ {
				path[i] = path[i] + spec[i]
			}

			dfs(price, special, needs, path, cur+spec[len(spec)-1], cost, k)

			// cancle choose
			for i := 0; i < len(price); i++ {
				path[i] = path[i] - spec[i]
			}
		}
	}
	// 如果还不够就选单品
	for k, v := range price {
		cur += (needs[k] - path[k]) * v
	}

	if cur < *cost {
		*cost = cur
	}
	return
}

func same(needs []int, path []int) bool {
	if len(needs) != len(path) {
		return false
	}
	for i := 0; i < len(needs); i++ {
		if needs[i] != path[i] {
			return false
		}
	}
	return true
}
