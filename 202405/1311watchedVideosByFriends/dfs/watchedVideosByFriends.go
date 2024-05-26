package main

import (
	"fmt"
	"sort"
)

func main() {
	ww := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	ff := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	fmt.Println(watchedVideosByFriends(ww, ff, 0, 1))
	fmt.Println(watchedVideosByFriends(ww, ff, 0, 2))

}

type video struct {
	K string
	V int
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	movies := make(map[string]int)
	visit := make(map[int]bool)
	dfs(watchedVideos, friends, id, level, movies, visit)
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

// todo 会出错，但是不啥出错，不知道原因
// 查找id 的好友观看的电影
func dfs(watchedVideos [][]string, friends [][]int, id int, lev int, mov map[string]int, visit map[int]bool) {
	if visit[id] {
		return
	}
	visit[id] = true
	if lev == 0 {
		for _, f := range watchedVideos[id] {
			mov[f]++
		}
		return
	}

	// 找他的朋友
	for _, f := range friends[id] {
		dfs(watchedVideos, friends, f, lev-1, mov, visit)
	}
}
