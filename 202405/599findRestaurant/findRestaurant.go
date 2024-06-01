package main

func main() {

}

func findRestaurant(list1 []string, list2 []string) []string {
	l1 := make(map[string]int)
	l2 := make(map[string]int)
	for i, k := range list1 {
		l1[k] = i + 1
	}
	for i, k := range list2 {
		l2[k] = i + 1
	}
	mx := max(len(list1)+len(list2)) + 10
	ans := make(map[string]int)

	for k, v := range l1 {
		if l2[k] > 0 {
			if v+l2[k] == mx {
				ans[k] = mx
			} else if v+l2[k] < mx {
				ans = make(map[string]int)
				mx = l2[k] + v
				ans[k] = mx
			}
		}
	}
	res := make([]string, 0)
	for k := range ans {
		res = append(res, k)
	}
	return res
}
