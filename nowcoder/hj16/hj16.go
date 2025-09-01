package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	first := strings.Fields(input.Text())
	var n, m int
	if len(first) >= 2 {
		n, _ = strconv.Atoi(first[0])
		m, _ = strconv.Atoi(first[1])
	} else {
		n, _ = strconv.Atoi(first[0])
		input.Scan()
		m, _ = strconv.Atoi(strings.TrimSpace(input.Text()))
	}
	nodes := make([]Node, 0)
	for i := 0; i < m; i++ {
		input.Scan()
		split := strings.Split(input.Text(), " ")
		price, _ := strconv.Atoi(split[0])
		wight, _ := strconv.Atoi(split[1])
		pa, _ := strconv.Atoi(split[2])
		nodes = append(nodes, Node{
			price: price,
			wight: wight,
			id:    i + 1,
			pa:    pa,
		})
	}
	groups := help(nodes)
	ans := sove(n, groups)
	fmt.Printf("%d", ans)
}

func help(node []Node) [][]Node {
	groups := make([][]Node, 0)
	mainItems := make(map[int]Node)
	accessories := make(map[int][]Node)

	// 分离主件和附件，并计算满意度
	for _, no := range node {
		no.va = no.price * no.wight // 满意度 = 价格 × 重要度
		if no.pa == 0 {
			mainItems[no.id] = no
		} else {
			accessories[no.pa] = append(accessories[no.pa], no)
		}
	}

	// 为每个主件生成所有可能的组合（最多 4 种：主件，主+附1，主+附2，主+附1+附2）
	for _, mainItem := range mainItems {
		options := make([]Node, 0, 4)
		// 只买主件
		options = append(options, Node{price: mainItem.price, va: mainItem.va})

		if accs, exists := accessories[mainItem.id]; exists {
			if len(accs) >= 1 {
				options = append(options, Node{price: mainItem.price + accs[0].price, va: mainItem.va + accs[0].va})
			}
			if len(accs) >= 2 {
				options = append(options, Node{price: mainItem.price + accs[1].price, va: mainItem.va + accs[1].va})
				options = append(options, Node{price: mainItem.price + accs[0].price + accs[1].price, va: mainItem.va + accs[0].va + accs[1].va})
			}
		}

		groups = append(groups, options)
	}

	return groups
}

type Node struct {
	price int
	wight int
	id    int
	pa    int
	va    int
}

func sove(num int, groups [][]Node) int {
	// 分组背包：每组（一个主件及其附件组合）最多选择一种
	dp := make([]int, num+1)

	for _, options := range groups {
		for j := num; j >= 0; j-- {
			best := dp[j] // 不买该组任何物品
			for _, opt := range options {
				if j >= opt.price {
					best = Max(best, dp[j-opt.price]+opt.va)
				}
			}
			dp[j] = best
		}
	}

	return dp[num]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
王强决定把年终奖用于购物，他把想买的物品分为两类：主件与附件。
∙

∙主件可以没有附件，至多有
2
2 个附件。附件不再有从属于自己的附件。
∙

∙若要购买某附件，必须先购买该附件所属的主件，且每件物品只能购买一次。

王强查到了
m 件物品的价格，而他只有
n 元的预算。为了先购买重要的物品，他给每件物品规定了一个重要度，用整数
1∼5 表示。他希望在不超过预算的前提下，使满意度最大。
满意度是 价格乘以重要度
*/
