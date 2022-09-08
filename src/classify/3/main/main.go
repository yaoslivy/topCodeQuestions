package main

func lengthOfLongestSubstring(s string) int {
	// 基于滑动窗口的思想，保持窗口内元素为不重复字符，当右边界下一字符为窗口内重复字符时，左边界向右收缩
	if len(s) <= 1 {
		return len(s)
	}
	mm := make(map[byte]bool)
	right := -1
	res := 1
	for i := 0; i < len(s)-1; i++ {
		//右边界向右扩，保持窗口内元素不重复
		for right+1 < len(s) && mm[s[right+1]] == false {
			mm[s[right+1]] = true
			right++
		}
		//到右边界
		res = max(res, right-i+1)
		//right+1为当前窗口内重复字符
		mm[s[i]] = false
	}
	return res
}

//暴力解法 超时
func lengthOfLongestSubstring2(s string) int {
	//直接暴力找到所有起始，终止的不含重复字符的最长子串
	if len(s) <= 1 {
		return len(s)
	}
	res := 1
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			//遍历所有情况，i为起点，j为终点，当这一段没有重复字符，则收集最长子串长度
			mm := make(map[byte]bool)
			flag := true
			for k := i; k <= j; k++ {
				if mm[s[k]] == true { //当前字符在前面出现过，则该段序列含重复字符
					flag = false
					break
				}
				mm[s[k]] = true
			}
			if flag == true {
				res = max(res, j-i+1)
			} else { //已经出现重复，则开始下一起始位置
				break
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
