package main

func nextGreaterElements(nums []int) []int {
	//使用单调栈记录之前访问过的值的下标，单调栈中从栈底到栈顶对于的元素值从大到小
	//当遍历到的元素值大于栈顶，则栈顶元素出栈，记录结果值
	S := make([]int, 0)
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1 //找不到则默认为-1
	}
	//遍历两次，即可将循环找出后一个更大的数，循环一遍都找不到则不存在比该栈顶元素大的值
	for i := 0; i < 2*len(nums); i++ {
		//当前元素比栈顶元素值大则一直出栈
		for len(S) != 0 && nums[i%len(nums)] > nums[S[len(S)-1]] {
			//记录当前栈顶元素下标的后一个最大值
			res[S[len(S)-1]] = nums[i%len(nums)]
			S = S[:len(S)-1]
		}
		//栈为空或当前元素值比栈顶元素值小
		S = append(S, i%len(nums))
	}
	return res
}
