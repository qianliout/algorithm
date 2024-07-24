package main

func main() {

}

// num 仅由数字组成且不含前导零
func largestOddNumber(num string) string {
	i := len(num) - 1
	for ; i >= 0; i-- {
		if (int(num[i])-int('0'))%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
