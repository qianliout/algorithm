package main

func main() {

}

type UndergroundSystem struct {
	IdStart   map[int]int
	IdStation map[int]string
	Gap       map[string]map[string][]int // start->end -time

}

func Constructor() UndergroundSystem {
	c := UndergroundSystem{
		IdStart:   map[int]int{},
		IdStation: map[int]string{},
		Gap:       map[string]map[string][]int{},
	}
	return c
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.IdStart[id] = t
	this.IdStation[id] = stationName

}

func (this *UndergroundSystem) CheckOut(id int, endStation string, endAt int) {
	startTime := this.IdStart[id]
	startState := this.IdStation[id]
	if this.Gap[startState] == nil {
		this.Gap[startState] = make(map[string][]int)
	}
	sub := endAt - startTime
	this.Gap[startState][endStation] = append(this.Gap[startState][endStation], sub)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	list := this.Gap[startStation][endStation]
	if len(list) == 0 {
		return 0
	}
	sum := 0
	for _, ch := range list {
		sum += ch
	}
	return float64(sum) / float64(len(list))
}
