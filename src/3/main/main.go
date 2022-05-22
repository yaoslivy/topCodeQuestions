package main

import "fmt"

func main() {
	mm := map[byte]int{}
	mm2 := make(map[byte]int, 0)
	fmt.Println(mm)
	fmt.Println(mm2)
	fmt.Println()
}

func lengthOfLongestSubstring(s string) int {
	mm := make(map[byte]int, 0) //记录字符出现的次数
	r := -1                     //确定i-r内无重复的字符区间
	res := 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(mm, s[i-1]) //每当i向右滑动一下时就在map中移除前一个字符
		}
		for r+1 < len(s) && mm[s[r+1]] == 0 { //在右边的的值没有出现过的情况下向右移动，否则以i为起点的最大字串长度就是r-i+1
			mm[s[r+1]]++
			r++
		}
		if res < r-i+1 { //确定每次i-r的最大值
			res = r - i + 1
		}
	}
	return res
}
