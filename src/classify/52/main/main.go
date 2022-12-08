package main

func totalNQueens(n int) int {
	//直接回溯统计所有可能的情况
	strs := ""
	for i := 0; i < n; i++ {
		strs += "."
	}
	res = make([][]string, 0)
	find(n, []string{}, strs)
	return len(res)
}

var res [][]string //记录所有可能情况 oneRes 记录一次摆放情况
func find(n int, oneRes []string, strs string) {
	if n == len(oneRes) {
		copyArr := make([]string, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}
	for i := 0; i < n; i++ {
		if isValid(oneRes, i) {
			newStrs := strs[:i] + "Q" + strs[i+1:]
			oneRes = append(oneRes, newStrs)
			find(n, oneRes, strs)
			oneRes = oneRes[:len(oneRes)-1]
		}
	}
}

//在已有行的下一行的col列上放置皇后，是否会被攻击
func isValid(oneRes []string, col int) bool {
	//第一行随便放置在哪一列
	if len(oneRes) == 0 {
		return true
	}
	//不能在同一列
	for i := 0; i < len(oneRes); i++ {
		if oneRes[i][col] == 'Q' {
			return false
		}
	}
	//不同在斜对角线上
	//不能在左斜线上
	for i, j := len(oneRes)-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if oneRes[i][j] == 'Q' {
			return false
		}
	}
	//不能在右斜线上
	for i, j := len(oneRes)-1, col+1; i >= 0 && j < len(oneRes[0]); i, j = i-1, j+1 {
		if oneRes[i][j] == 'Q' {
			return false
		}
	}
	return true
}
