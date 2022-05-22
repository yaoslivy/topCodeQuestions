package main

import "fmt"

func main() {
	n := 4
	k := 2
	ans := combine(n, k)
	fmt.Println(ans)
}

var res [][]int

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}
	res = make([][]int, 0)
	find([]int{}, n, k, 1)
	return res
}

func find(path []int, n int, k int, startIndex int) {
	if len(path) == k { //找到结果
		copyArr := make([]int, k)
		copy(copyArr, path)
		res = append(res, copyArr)
		return
	}
	if startIndex > n-(k-len(path))+1 { //k-len(path)：还需要的元素个数，
		return
	}

	for i := startIndex; i <= n; i++ { //横向遍历，不能取比startIndex小的数
		path = append(path, i)    //处理当前节点
		find(path, n, k, i+1)     // 递归
		path = path[:len(path)-1] // 回溯
	}
}
