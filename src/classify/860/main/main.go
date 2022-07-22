package main

func lemonadeChange(bills []int) bool {
	// 给10的找5，给20的优先找10+5，再找5*3，如果无法找则返回false
	arr := make([]int, 2) //0 0  代表5 10 的个数
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			arr[0]++
		case 10:
			if arr[0] < 1 {
				return false
			} else {
				arr[0]--
				arr[1]++
			}
		case 20:
			if arr[1] > 0 && arr[0] > 0 {
				arr[1]--
				arr[0]--
			} else if arr[0] > 2 {
				arr[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func lemonadeChange2(bills []int) bool {
	mm := make(map[int]int)
	if bills[0] != 5 {
		return false
	}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			mm[5]++
		} else if bills[i] == 10 {
			if mm[5] <= 0 {
				return false
			}
			mm[10]++
			mm[5]--
		} else if bills[i] == 20 {
			if mm[5] <= 0 { //一定要有5元的
				return false
			}
			if mm[10] >= 1 {
				mm[10]--
				mm[5]--
				mm[20]++
				continue
			}
			if mm[5] >= 3 {
				mm[5] -= 3
				mm[20]++
			} else {
				return false
			}

		}
	}
	return true
}
