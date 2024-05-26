package main

import (
	"fmt"
	"sort"
)

func main() {
	ww := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	ff := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	// fmt.Println(watchedVideosByFriends(ww, ff, 0, 1))
	fmt.Println(watchedVideosByFriends(ww, ff, 0, 2))
}

type video struct {
	K string
	V int
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	queue := make([]int, 0)
	queue = append(queue, id)
	visit := make(map[int]bool)
	visit[id] = true
	for len(queue) > 0 {
		if level == 0 {
			break
		}
		level--
		lev := make([]int, 0)
		for _, no := range queue {
			for _, f := range friends[no] {
				if visit[f] {
					continue
				}
				visit[f] = true
				lev = append(lev, f)
			}
		}
		queue = lev
	}
	movies := make(map[string]int)
	for _, no := range queue {
		for _, f := range watchedVideos[no] {
			movies[f]++
		}
	}

	ans := make([]video, 0)
	for i, v := range movies {
		if v > 0 {
			ans = append(ans, video{K: i, V: v})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].V < ans[j].V {
			return true
		} else if ans[i].V > ans[j].V {
			return false
		}
		return ans[i].K < ans[j].K
	})

	res := make([]string, 0)
	for i := range ans {
		res = append(res, ans[i].K)
	}
	return res
}
