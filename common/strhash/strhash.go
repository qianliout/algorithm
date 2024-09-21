package strhash

import (
	"math"
)

/*
字符串编码，类试字符串的前缀和,可以快速找一个长字符串中相同的 substring
*/

type StrHash struct {
	Str  string
	Hash []int64
	P    []int64
	Base int64
	Mod  int64
}

// NewStrHash 计算字符串的哈希值和幂基数数组。
// 该函数使用特定的基数和模数计算字符串的滚动哈希值，
// 以便于在后续进行字符串匹配或哈希比较时提高效率。
// 参数 str 是需要计算哈希值的字符串。
// 参数 base 是计算哈希值时使用的基数。如果全是小写 base 就可以设置为27
// 返回值 StrHash 是一个结构体，包含哈希值数组和幂基数数组。
func NewStrHash(str string, base int64) *StrHash {
	// n 为字符串的长度
	n := len(str)
	// hash 数组用于存储每个前缀的哈希值
	hash := make([]int64, n+1)
	// p 数组用于存储基数的幂次方值
	p := make([]int64, n+1)
	// 初始化 p[0] 为 1，因为任何数的 0 次方都是 1,也表示空字符串的 hash值是0
	p[0] = 1

	// mod 为取模操作的模数，用于防止整数溢出
	mod := int64(math.Pow10(9)) + 7

	// 遍历字符串，计算每个前缀的哈希值和幂基数数组
	for i, ch := range str {
		// 计算当前前缀的哈希值，使用之前的哈希值乘以基数加上当前字符编码，然后取模
		hash[i+1] = (hash[i]*base + encode(byte(ch))) % mod
		// 计算当前前缀的幂基数，即基数的 i 次方取模
		p[i+1] = p[i] * base % mod
	}
	// 返回计算好的哈希值数组和幂基数数组的结构体指针
	return &StrHash{Str: str, Hash: hash, Base: base, Mod: mod, P: p}
}

func (s *StrHash) Get(le, ri int) int64 {
	return (s.Hash[ri] - s.Hash[le]*s.P[ri-le]%s.Mod + s.Mod) % s.Mod
}

func encode(ch byte) int64 {
	return int64(ch) - int64('a') + 1
}
