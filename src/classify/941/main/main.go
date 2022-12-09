package main

func validMountainArray(arr []int) bool {
	// return validMountainArrayTraverse(arr)
	return validMountainArrayDoublePointer(arr)
}

//双指针思路
func validMountainArrayDoublePointer(arr []int) bool {
	//序列严格先递增，再递减，设置双指针向中间逼近
	//如果双指针不能相遇，则为false
	if len(arr) < 3 {
		return false
	}
	left, right := 0, len(arr)-1
	//找到递增右边界
	for left < len(arr)-1 && arr[left] < arr[left+1] {
		left++
	}
	//找到递减左边界
	for right > 0 && arr[right] < arr[right-1] {
		right--
	}
	//两边界重合，且不是在序列边界处
	if left == right && left != 0 && right != len(arr)-1 {
		return true
	}
	return false
}

//观察相邻变化趋势
func validMountainArrayTraverse(arr []int) bool {
	//长度不符合
	if len(arr) < 3 {
		return false
	}
	preDiff := arr[1] - arr[0]
	cnt := 0 //记录相邻差值正负变化的次数
	for i := 2; i < len(arr); i++ {
		curDiff := arr[i] - arr[i-1]
		//开始递增必须
		if preDiff <= 0 {
			return false
		}
		//相邻不能有相等的值
		if curDiff == 0 {
			return false
		}
		//已存在一个上下坡，不能再有另一个上下坡
		if cnt == 1 && curDiff > 0 {
			return false
		}
		if cnt == 0 && curDiff < 0 {
			cnt++
		}
	}
	if cnt == 1 {
		return true
	}
	return false
}
