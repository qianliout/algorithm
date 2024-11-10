package main

func main() {

}

func minimumFinishTime1(tires [][]int, changeTime int, numLaps int) int {
	// 根据题目的数据范围一个轮胎最多跑17圈，所以设置一个上限17
	// 首先预处理出连续使用同一个轮胎跑 x 圈的最小耗时，记作 minSec[x]，这可以通过遍历每个轮胎计算出来。
	minSec := make([]int, 18)
	inf := 1 << 30
	for i := range minSec {
		minSec[i] = inf
	}
	for _, ch := range tires {
		f, r := ch[0], ch[1]
		x, ti, sum := 1, f, 0
		for ti <= changeTime+f {
			sum += ti
			minSec[x] = min(minSec[x], sum)
			ti = ti * r
			x++
		}
	}
	f := make([]int, numLaps+1)
	f[0] = -changeTime // 这里怎么想
	for i := 1; i <= numLaps; i++ {
		f[i] = inf
		for j := 1; j <= 17 && j <= i; j++ {
			f[i] = min(f[i], f[i-j]+minSec[j])
		}
		f[i] += changeTime
	}

	return f[numLaps]
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	// 根据题目的数据范围一个轮胎最多跑17圈，所以设置一个上限17
	// 首先预处理出连续使用同一个轮胎跑 x 圈的最小耗时，记作 minSec[x]，这可以通过遍历每个轮胎计算出来。
	// minSec[i] 表示只使用一种轮胎（不知道是那种轮胎），最少需要多少秒
	minSec := make([]int, 18)
	inf := 1 << 30
	for i := range minSec {
		minSec[i] = inf
	}
	for _, ch := range tires {
		f, r := ch[0], ch[1]
		x, ti, sum := 1, f, 0
		for ti <= changeTime+f {
			sum += ti
			minSec[x] = min(minSec[x], sum)
			ti = ti * r
			x++
		}
	}
	f := make([]int, numLaps+1)
	f[0] = -changeTime // 这里怎么想,这里只能赋这个值

	for i := 1; i <= numLaps; i++ {
		f[i] = inf
		for j := 1; j <= 17 && j <= i; j++ {
			// 当i==j时，也就是从第0个开始里，这里还是加了 changeTime 所以初值得用-changeTime
			f[i] = min(f[i], changeTime+f[i-j]+minSec[j])
		}
	}

	return f[numLaps]
}
