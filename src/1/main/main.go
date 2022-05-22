package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(arr, target))
}

func twoSum(nums []int, target int) []int {
	mm := make(map[int]int, 0)
	for i, val := range nums { //每遍历一个数都当哈希表中查找是否存在和哈希表中的值相加=target的情况,不存在则加入哈希表中
		if index, ok := mm[target-val]; ok {
			return []int{index, i}
		}
		mm[val] = i
	}
	return nil
}
