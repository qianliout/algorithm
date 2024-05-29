package main

import (
	"fmt"
)

func main() {
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("1/3-1/2"))
	fmt.Println(fractionAddition("-7/3"))
}

type pair struct{ x, y int }

func fractionAddition(expression string) string {
	// expression = "-1/2+1/2+1/3"
	op, num, fir, sec := 1, 0, 0, 0
	stark := make([]pair, 0)
	for _, ch := range expression {
		switch ch {
		case ' ':
			continue
		case '+':
			sec = num
			op, num = 1, 0
			if sec == 0 {
				continue
			}
			stark = append(stark, pair{fir, sec})
		case '-':
			op = -1
			sec = num
			num = 0
			if sec == 0 {
				continue
			}
			stark = append(stark, pair{fir, sec})
		case '/':
			fir = num * op
			// +3/-3 // 这种情部分就会出错
			op = 1
			num = 0

		default:
			num = num*10 + int(ch) - int('0')
		}
	}
	if num > 0 {
		stark = append(stark, pair{fir * op, num})
	}

	for len(stark) > 2 {
		a, b := stark[0], stark[1]
		stark = stark[2:]
		stark = append(stark, add(a, b))
	}
	if len(stark) == 1 {
		return pt(stark[0])
	}
	c := add(stark[0], stark[1])
	return pt(c)
}

func pt(p pair) string {
	if p.y < 0 {
		return fmt.Sprintf("%d/%d", -p.x, -p.y)
	}
	if p.y > 0 {
		return fmt.Sprintf("%d/%d", p.x, p.y)
	}
	return ""
}

func add(fir, sec pair) pair {
	a := fir.x*sec.y + fir.y*sec.x
	b := fir.y * sec.y
	// 化简单
	c := gcb(a, b)
	return pair{a / c, b / c}
}

// 求最大公约数
func gcb(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcb(b, a%b)
}
