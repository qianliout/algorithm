package main

func main() {

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
	return dfs(price, special, needs)
}

func dfs(price []int, special [][]int, needs []int) int {
	minPrice := 0
	for i, ch := range needs {
		minPrice += price[i] * ch
	}
loop:
	for _, ch := range special {
		newNeeds := make([]int, len(needs))
		copy(newNeeds, needs)
		for i := range newNeeds {
			newNeeds[i] -= ch[i]
			if newNeeds[i] < 0 {
				continue loop
			}
		}
		curPrice := dfs(price, special, newNeeds) + ch[len(price)]
		minPrice = min(minPrice, curPrice)
	}
	return minPrice
}
