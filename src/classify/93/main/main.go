package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "25525511135" //["255.255.11.135","255.255.111.35"]
	//s := "101023" //["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
	fmt.Println(restoreIpAddresses(s))
}

var res []string

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	res = make([]string, 0)
	find(s, "", 0, 0)
	return res
}

func find(s string, oneRes string, nums int, start int) {
	if nums == 3 { //如果字符串添加到3个.，则判断第四个序列是否合理
		if isValid(s[len(oneRes)-3:]) {
			res = append(res, oneRes+s[len(oneRes)-3:])
		}
		return
	}

	for i := start; i < len(s); i++ {
		if isValid(s[start : i+1]) { //截取时需要注意，i+1位置不会取
			oneResLen := len(oneRes)
			oneRes = oneRes + s[start:i+1] + "."
			find(s, oneRes, nums+1, i+1)
			oneRes = oneRes[:oneResLen]
		}
	}
}

func isValid(s string) bool {
	//长度
	if len(s) <= 0 || len(s) > 3 {
		return false
	}
	//数值
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	nums, _ := strconv.Atoi(s)
	if nums < 0 || nums > 255 {
		return false
	}
	//第一个位置为0，第二个位置不能有值
	if s[0] == '0' && len(s) > 1 {
		return false
	}

	return true
}
