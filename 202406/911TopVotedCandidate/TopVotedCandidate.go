package main

func main() {

}

type TopVotedCandidate struct {
	Data map[int][]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	data := make(map[int][]int)
	n := len(times)
	for i := 0; i < n; i++ {
		data[persons[i]] = append(data[persons[i]], times[i])
	}
	return TopVotedCandidate{Data: data}

}

func (this *TopVotedCandidate) Q(t int) int {

}
