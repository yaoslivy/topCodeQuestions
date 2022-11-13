package main

func trap(height []int) int {
	// return trapByStack(height)
	// return trapDoublePointer(height)
	return trapByDynamic(height)
}

//动态规划思路
func trapByDynamic(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	//第一个位置和最后一个位置不可能收集雨水，收集中间每个位置左右两边的最高值，动态规划
	// leftMaxH[i] = max(height[i], leftMaxH[i-1])
	// rightMaxH[i] = max(height[i], rightMaxH[i+1])
	// leftMaxH[0] = height[0] rightMaxH[len(height)-1] = height[len(height)-1]
	leftMaxH := make([]int, len(height))
	rightMaxH := make([]int, len(height))
	leftMaxH[0] = height[0]
	rightMaxH[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height)-1; i++ {
		leftMaxH[i] = max(height[i], leftMaxH[i-1])
	}
	for j := len(height) - 2; j >= 1; j-- {
		rightMaxH[j] = max(height[j], rightMaxH[j+1])
	}
	//计算中间柱子收集雨水量
	res := 0
	for i := 1; i < len(height)-1; i++ {
		h := min(leftMaxH[i], rightMaxH[i]) - height[i]
		if h > 0 {
			res += h * 1
		}
	}
	return res
}

//双指针思路
func trapByDouble(height []int) int {
	//统计每一个位置接的雨水，如果当前位置可以接雨水，则左右两边柱子高度有 > 当前柱子高度的情况则可以接雨水
	maxLeft := make([]int, len(height))  //i位置柱子左边最大的柱子下标
	maxRight := make([]int, len(height)) //i位置柱子右边最大的柱子下标
	maxLeft[0] = 0
	maxRight[len(height)-1] = len(height) - 1
	for i := 1; i < len(height); i++ {
		if height[i] > height[maxLeft[i-1]] {
			maxLeft[i] = i
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > height[maxRight[i+1]] {
			maxRight[i] = i
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}
	//首尾不可能接雨水
	res := 0
	for i := 1; i < len(height)-1; i++ {
		h := min(height[maxRight[i]], height[maxLeft[i]]) - height[i]
		if h <= 0 {
			continue
		}
		res += h
	}
	return res
}

//双指针思路
func trapDoublePointer(height []int) int {
	//遍历所有下标，以下标为起点，向右找到右侧最高的柱子，向左找到左侧最高的柱子，根据高度差判断是否当前柱子可以收集雨水
	if len(height) <= 2 {
		return 0
	}
	res := 0
	//所有可能收集雨水的下标，第一个和最后一个下标不可能收集
	for i := 1; i < len(height)-1; i++ {
		leftH, rightH := height[i], height[i]
		//向右拓展找到右侧最高的柱子
		for right := i + 1; right < len(height); right++ {
			rightH = max(rightH, height[right])
		}
		//向左拓展找到左侧最高的柱子
		for left := i - 1; left >= 0; left-- {
			leftH = max(leftH, height[left])
		}
		//计算当前柱子收集的雨水
		h := min(leftH, rightH) - height[i]
		if h > 0 {
			res += h * 1
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//单调栈思路
func trapByStack(height []int) int {
	//使用单调栈记录访问过的高度的下标，保持栈从栈底到栈顶从高度值从大到小
	//如果当前元素值<=栈顶元素值小，直接入栈，小于则找到凹槽处，计算雨水
	if len(height) <= 2 {
		return 0
	}
	S := make([]int, 0)
	res := 0
	for i := 0; i < len(height); i++ {
		//栈不为空且当前元素大于栈顶元素则计算
		for len(S) != 0 && height[i] > height[S[len(S)-1]] {
			//当前元素和栈顶两个元素构成一个凹槽
			mid := S[len(S)-1]
			S = S[:len(S)-1]
			if len(S) == 0 {
				break
			}
			h := min(height[i], height[S[len(S)-1]]) - height[mid]
			w := i - S[len(S)-1] - 1 //中间高度
			res += h * w
		}
		//栈为空或当前元素小于栈顶元素则入栈
		S = append(S, i)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
