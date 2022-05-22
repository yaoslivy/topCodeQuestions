package main

func lemonadeChange(bills []int) bool {
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
