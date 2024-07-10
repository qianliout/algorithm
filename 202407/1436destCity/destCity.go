package main

func main() {

}

func destCity(paths [][]string) string {
	road := make(map[string]string)
	for _, ch := range paths {
		x, y := ch[0], ch[1]
		road[x] = y
	}
	ans := paths[0][0]
	for {
		if road[ans] != "" {
			ans = road[ans]
		} else {
			// 题目保证了有解
			return ans
		}
	}
}
