package main

func main() {

}

// 有一个餐厅，只有一位厨师。你有一个顾客数组 customers ，其中 customers[i] = [arrivali, timei] ：
//    arrivali 是第 i 位顾客到达的时间，到达时间按 非递减 顺序排列。
//    timei 是给第 i 位顾客做菜需要的时间。

func averageWaitingTime(customers [][]int) float64 {
	cnt := 0
	preEnd := 0
	for _, ch := range customers {

		arrive, used := ch[0], ch[1]

		start := max(preEnd, arrive)
		end := start + used
		cnt += end - arrive
		preEnd = end
	}
	return float64(cnt) / float64(len(customers))
}
