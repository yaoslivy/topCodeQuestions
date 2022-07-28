package main

func generateMatrix(n int) [][]int {
	//设置上右下左边界值，赋值一层后就改变边界值
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for num <= n*n {
		//最上面一行赋值
		for j := left; j <= right; j++ {
			res[top][j] = num
			num++
		}
		top++ //该行已经赋值
		//最右边一列赋值
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right-- //该列已经赋值
		//最底边一行赋值
		for j := right; j >= left; j-- {
			res[bottom][j] = num
			num++
		}
		bottom-- //该行已经赋值
		//最左边一列赋值
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++ //该列已经赋值
	}
	return res
}
