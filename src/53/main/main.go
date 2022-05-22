package main

import "fmt"

func main() {
	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray3(arr))
}

type Status struct {
	totalSum int //表示[l, r]区间内所有值的和
	leftSum  int //表示[l, r]内以l为左端点的最大子段和
	rightSum int //表示[l, r]内以r为右端点的最大子段和
	maxSum   int //表示[l, r]内最大子段和
}

func maxSubArray3(nums []int) int {
	return mergeSeq(nums, 0, len(nums)-1).maxSum
}

func mergeSeq(arr []int, left, right int) Status {
	if left == right { //若当前arr中只有一个元素，则拆解到了最小子序列
		return Status{arr[left], arr[left], arr[left], arr[left]}
	}
	mid := (left + right) >> 1
	leftSeq := mergeSeq(arr, left, mid)
	rightSeq := mergeSeq(arr, mid+1, right)
	return merge(leftSeq, rightSeq)
}

func merge(leftSeq, rightSeq Status) Status {
	totalSum := leftSeq.totalSum + rightSeq.totalSum                                       //合并后的元素总和=左序列总和+右序列总和
	leftSum := max(leftSeq.leftSum, leftSeq.totalSum+rightSeq.leftSum)                     //以左端点的最大子段和=左序列的以左端点的最大子段和  / 左序列的总和+右序列的以左端点的最大子段和
	rightSum := max(rightSeq.rightSum, leftSeq.rightSum+rightSeq.totalSum)                 //以右端点的最大子段和= 右序列的以右端点的最大子段和 / 左序列的的以右端点的最大子段和+右序列的总和
	maxSum := max(max(leftSeq.maxSum, rightSeq.maxSum), leftSeq.rightSum+rightSeq.leftSum) //合并序列的最大子段和= 左序列的最大子段和 / 右序列的最大子段和 / 左序列的以右端点的最大子段和+右序列的以左端点的最大子段和
	return Status{totalSum: totalSum, leftSum: leftSum, rightSum: rightSum, maxSum: maxSum}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1) // 每一个位置记录以i为结尾序列的最大和的连续子数组
	res := nums[0]
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		if dp[i]+nums[i] > nums[i] {
			dp[i+1] = dp[i] + nums[i]
		} else {
			dp[i+1] = nums[i]
		}
		if res < dp[i+1] {
			res = dp[i+1]
		}
	}
	return res
}

func maxSubArray2(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] { //滚动数组，每个位置都记录了以i位置结尾的序列最大值
			nums[i] += nums[i-1]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}
