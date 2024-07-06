package main

import (
	"sort"
)

func main() {

}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	ans := make([]Restaurant, 0)
	for _, ch := range restaurants {
		res := Restaurant{
			ID:    ch[0],
			Rat:   ch[1],
			Veg:   ch[2],
			Price: ch[3],
			Dis:   ch[4],
		}
		if veganFriendly == 1 && res.Veg == 1 {
			continue
		}
		if res.Price > maxPrice {
			continue
		}
		if res.Dis > maxDistance {
			continue
		}
		ans = append(ans, res)
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].Rat > ans[j].Rat {
			return true
		} else if ans[i].Rat < ans[j].Rat {
			return false
		} else {
			return ans[i].ID < ans[j].ID
		}
	})
	res := make([]int, 0)
	for _, ch := range ans {
		res = append(res, ch.ID)
	}
	return res
}

// 给你一个餐馆信息数组 restaurants，其中  restaurants[i] = [idi, ratingi, veganFriendlyi, pricei, distancei]。你必须使用以下三个过滤器来过滤这些餐馆信息。

type Restaurant struct {
	ID    int
	Rat   int
	Veg   int
	Price int
	Dis   int
}
