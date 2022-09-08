package main

func largestRectangleArea(heights []int) int {
	// return largestRectangleAreaStack(heights)
	// return largestRectangleAreaDoublePointer(heights)
	return largestRectangleAreaDynamic(heights)
}

//动态规划
func largestRectangleAreaDynamic(heights []int) int {
	//使用两个数组分别记录当前位置左右两边第一个小于当前位置的下标，而不是左边第一个小于该柱子的高度
	minLeftIndex := make([]int, len(heights))
	minRightIndex := make([]int, len(heights))
	//记录每个柱子，左边第一个小于该柱子的下标
	minLeftIndex[0] = -1 //第一个位置
	for i := 1; i < len(heights); i++ {
		left := i - 1
		for left >= 0 && heights[i] <= heights[left] {
			left = minLeftIndex[left]
		}
		minLeftIndex[i] = left
	}
	//记录每一个柱子，右边第一个小于该柱子的下标
	minRightIndex[len(heights)-1] = len(heights)
	for i := len(heights) - 2; i >= 0; i-- {
		right := i + 1
		for right < len(heights) && heights[i] <= heights[right] {
			right = minRightIndex[right]
		}
		minRightIndex[i] = right
	}
	//计算每一个柱子的为中心向两边拓展的面积
	res := 0
	for i := 0; i < len(heights); i++ {
		w := minRightIndex[i] - minLeftIndex[i] - 1
		res = max(res, w*heights[i])
	}
	return res
}

//使用双指针 O(n^2) 超时
func largestRectangleAreaDoublePointer(heights []int) int {
	//以每一个位置往左右两边扩展，左右两边都拓展到比当前位置更小的位置，再计算宽度和高度
	res := 0
	//每一个计算的位置
	for i := 0; i < len(heights); i++ {
		//找到左，右两边更小的位置
		left, right := i, i
		for right < len(heights) && heights[i] <= heights[right] {
			right++
		}
		for left >= 0 && heights[i] <= heights[left] {
			left--
		}
		//计算面积
		w := right - left - 1
		res = max(res, w*heights[i])
	}
	return res
}

//使用单调栈
func largestRectangleAreaStack(heights []int) int {
	//使用单调栈记录访问过的元素，栈中记录下标，栈底到栈顶元素值从小到大
	//若当前位置元素值小于栈顶元素值，则开始计算最大面积
	if len(heights) == 0 {
		return 0
	}
	S := make([]int, 0)
	res := 0
	//现在柱子左右两端都加个0，便于计算左右两端点覆盖的情况
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		//当前元素小于栈顶元素时计算面积
		for len(S) != 0 && heights[i] < heights[S[len(S)-1]] {
			mid := S[len(S)-1]
			S = S[:len(S)-1]
			if len(S) == 0 {
				break
			}
			w := i - S[len(S)-1] - 1
			h := heights[mid] // i值为2时 栈顶为6 w=1 h=6 退栈，栈顶为5 w=2 h=5
			res = max(res, w*h)
		}
		//栈为空或当前元素>=栈顶元素时入栈
		S = append(S, i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
