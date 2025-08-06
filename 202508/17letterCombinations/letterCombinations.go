package main

func main() {

}

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
	ss := []byte(digits)
	n := len(ss)
	var dfs func(i int, path []byte)
	dfs = func(i int, path []byte) {
		if i >= n {
			if len(path) > 0 {
				ans = append(ans, string(path))
			}
			return
		}
		kk := digitMap[ss[i]]
		for _, ch := range kk {
			path = append(path, ch)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []byte{})
	return ans
}
