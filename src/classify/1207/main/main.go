package main

import "sort"

func uniqueOccurrences(arr []int) bool {
	//一次遍历map统计不同数之间的出现个数，diffArr记录不同数
	arrMap := make(map[int]int)
	arrDiff := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if _, ok := arrMap[arr[i]]; !ok {
			arrDiff = append(arrDiff, arr[i])
		}
		arrMap[arr[i]]++
	}
	//再使用一个map记录不同的出现次数
	diffMap := make(map[int]bool)
	for i := 0; i < len(arrDiff); i++ {
		if diffMap[arrMap[arrDiff[i]]] {
			return false
		}
		diffMap[arrMap[arrDiff[i]]] = true
	}
	return true
}

func uniqueOccurrences2(arr []int) bool {
	//一次遍历map统计不同数之间的出现个数，diffArr记录不同数
	arrMap := make(map[int]int)
	arrDiff := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if _, ok := arrMap[arr[i]]; !ok {
			arrDiff = append(arrDiff, arr[i])
		}
		arrMap[arr[i]]++
	}
	//按出现的次数排序
	sort.Slice(arrDiff, func(i, j int) bool {
		return arrMap[arrDiff[i]] < arrMap[arrDiff[j]]
	})
	//一次遍历，观察相邻的数的出现次数是否相同
	for i := 0; i < len(arrDiff)-1; i++ {
		if arrMap[arrDiff[i]] == arrMap[arrDiff[i+1]] {
			return false
		}
	}
	return true
}
