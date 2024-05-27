package main

func main() {

}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	in := make(map[string][]string, len(recipes))
	deg := make(map[string]int)
	for i, r := range recipes {
		for _, ing := range ingredients[i] {
			in[ing] = append(in[ing], r)
		}

		deg[r] = len(ingredients[i])
	}
	ans := make([]string, 0)
	for len(supplies) > 0 {
		no := supplies[0]
		supplies = supplies[1:]
		for _, r := range in[no] {
			deg[r]--
			if deg[r] == 0 {
				ans = append(ans, r)
				supplies = append(supplies, r)
			}
		}
	}
	return ans
}
