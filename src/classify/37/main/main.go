package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	for _, val := range board {
		fmt.Printf("%c\n", val)
	}
	solveSudoku(board)
	fmt.Println("---------------")
	for _, val := range board {
		fmt.Printf("%c\n", val)
	}
}

func solveSudoku(board [][]byte) {
	find(board)
}

func find(board [][]byte) bool {
	//对每一个位置上逐步填充
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == '.' { //需要填充的位置
				for _, val := range "123456789" { //尝试填充
					if isValid(board, row, col, byte(val)) {
						board[row][col] = byte(val)
						if find(board) == true { //找到合适的一组立刻返回
							return true
						}
						board[row][col] = '.' //如果在某个位置发现前面填充的值有误则回溯重新尝试其他值
					}
				}
				return false //在需要填充的位置尝试了所有值都失败了
			}
		}
	}
	return true //所有位置填充完成
}

func isValid(board [][]byte, row, col int, val byte) bool {
	//判断同一行,和同一列
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] == val {
			return false
		}
		if board[i][col] == val {
			return false
		}
	}
	//判断子宫格 注意开始行和开始列的处理
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
