package main

func main() {}

func letterCombinations(digits string) []string {
	digitMap := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	ans := make([]string, 0)
	n := len(digits)
	path := make([]byte, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range digitMap[digits[i]] {
			path = append(path, c)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
