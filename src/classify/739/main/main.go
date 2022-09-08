package main

func dailyTemperatures(temperatures []int) []int {
	// return dailyTemperaturesTraverse(temperatures)
	return dailyTemperaturesByStack(temperatures)
}

//单调栈解法
func dailyTemperaturesByStack(temperatures []int) []int {
	//使用一个栈记录来记录每一个右边第一个比当前元素高的元素
	//栈记录元素下标，从栈顶到栈低保持保持递增
	//因为遍历到后面的元素若大于栈顶元素，则表明是栈顶元素的之后高的元素，一直出栈，并记录差值。
	S := make([]int, 0)
	res := make([]int, len(temperatures)) //后面没有更大的值则默认为0
	for i := 0; i < len(temperatures); i++ {
		//当前元素比栈顶元素值大，则说明时栈顶元素的后面更大的元素
		for len(S) != 0 && temperatures[i] > temperatures[S[len(S)-1]] {
			res[S[len(S)-1]] = i - S[len(S)-1]
			S = S[:len(S)-1]
		}
		//栈为空或当前元素比栈顶元素值小
		S = append(S, i)
	}
	return res
}

//暴力解法
func dailyTemperaturesTraverse(temperatures []int) []int {
	res := make([]int, len(temperatures))
	//需要确定的位置
	for i := 0; i < len(temperatures)-1; i++ {
		//下一个更高温度
		j := i + 1
		//找到位置j的温度>当前i位置的
		for j < len(temperatures) && temperatures[i] >= temperatures[j] {
			j++
		}
		if j == len(temperatures) { //没有比当前i位置高的
			res[i] = 0
		} else {
			res[i] = j - i
		}
	}
	return res
}
