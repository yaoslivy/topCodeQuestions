package main

import "fmt"

func main() {

	ans := solveNQueens(4)
	for _, val := range ans {
		fmt.Println(val)
	}

}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	strs := ""
	for i := 0; i < n; i++ {
		strs += "."
	}
	find(n, []string{}, strs)
	return res
}

func find(n int, oneRes []string, strs string) {
	if n == len(oneRes) { //若一次递归集合中满足条件
		copyArr := make([]string, n)
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}

	for i := 0; i < n; i++ { //尝试在同一行的不同列中放入
		if isValid(oneRes, i) == false { //如果在当前列放入不满足条件，则放入下一列
			continue
		}
		newStrs := strs[:i] + "Q" + strs[i+1:] //注意要新建一个变量
		oneRes = append(oneRes, newStrs)
		find(n, oneRes, strs)
		oneRes = oneRes[:len(oneRes)-1]
	}
}

func isValid(oneRes []string, col int) bool {
	if len(oneRes) == 0 { //当前放入的是第一行，一定可以放入
		return true
	}
	//判断列
	for _, val := range oneRes {
		if val[col] == 'Q' { //如果存在一行的列元素位置有Q则失败
			return false
		}
	}
	//判断左斜线上
	for r, c := len(oneRes)-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if oneRes[r][c] == 'Q' {
			return false
		}
	}
	//判断右斜线上
	for r, c := len(oneRes)-1, col+1; r >= 0 && c < len(oneRes[0]); r, c = r-1, c+1 {
		if oneRes[r][c] == 'Q' {
			return false
		}
	}
	return true
}
