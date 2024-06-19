package main

import "sort"

func main() {

}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]pair, 0)
	for i := range position {
		cars = append(cars, pair{pos: position[i], speed: speed[i]})
	}

	sort.Slice(cars, func(i, j int) bool { return cars[i].pos <= cars[j].pos })

	stark := make([]float64, 0)
	n := len(position)
	times := make([]float64, n)
	for i := range cars {
		times[i] = float64(target-cars[i].pos) / float64(cars[i].speed)
	}
	for _, ti := range times {
		// 这个做法太妙了，要学会才行
		for len(stark) > 0 && ti >= stark[len(stark)-1] {
			stark = stark[:len(stark)-1]
		}
		stark = append(stark, ti)
	}

	return len(stark)
}

type pair struct {
	pos   int // 启点
	speed int // 速度
}
