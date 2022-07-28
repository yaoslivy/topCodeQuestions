package main

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	resLen := 0
	//将前面等于val的都交换到后面去
	for i <= j { //包含一个元素情况
		if nums[i] == val {
			for i < j && nums[j] == val { //从后往前找到一个位置元素不为val
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
			if i != j { //不是相同的位置交换才将不为val个数++
				resLen++
			}
		} else { //不为val个数++
			resLen++
		}
		i++
	}
	return resLen
}
