package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumPerimeter(1))
}

func minimumPerimeter(neededApples int64) int64 {
	// 一个边长是 n,的花园，所有的苹果数是，2*n*(n+1)*(2n+1)
	// 求x向上的最大边长
	//  二分求左端点
	// 这里 ri 不能写成 math.Max,因为下在的乘法会越界
	le, ri := 0, int(math.Pow10(5))+10
	for le < ri {
		mid := le + (ri-le)/2
		if int64(2*mid*(mid+1)*(2*mid+1)) >= neededApples {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	// 返回的结果是所有的边长
	return int64(le) * 8
}
