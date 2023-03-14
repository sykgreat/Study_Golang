package DFS

import (
	"strconv"
	"testing"
)

// 假设我们有一个数字到字母表的映射:
// 1 ->['a', 'b ', 'c']
// 2 ->['d', 'e']
// 3 ->['f', 'g'， 'h']
// 实现一个函数，对于给定的一串数字，例如“1"、"233"，返回一个包含所有可能的组合的字符串列表。

var dict = map[int][]string{
	1: {"a", "b", "c"},
	2: {"d", "e"},
	3: {"f", "g", "h"},
}

func Test_DFS_1(t *testing.T) {
	var res []string
	letterCombinations("", "123", &res)
	t.Log(res)
}

func letterCombinations(str, num string, res *[]string) {
	if len(num) == 0 {
		*res = append(*res, str)
		return
	}

	i, err := strconv.ParseInt(num[:1], 10, 64)
	if err != nil {
		panic(err)
	}
	for _, v := range dict[int(i)] {
		letterCombinations(str+v, num[1:], res)
	}

	return
}
