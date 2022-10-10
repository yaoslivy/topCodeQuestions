package main

import "strings"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//使用一个map记录wordList中出现的所有字符串，便于快速查找
	wordListMap := make(map[string]bool)
	for _, val := range wordList {
		wordListMap[val] = true
	}
	if wordListMap[endWord] != true { //终止字符串不在转换列表中
		return 0
	}
	//记录是否访问过
	visitedPath := make(map[string]int) //查询到这个word的路径长度
	//初始化访问队列
	Q := make([]string, 0)
	Q = append(Q, beginWord)
	visitedPath[beginWord] = 1
	for len(Q) != 0 {
		curWord := Q[0]
		Q = Q[1:]
		//获取当前单词的路径长度
		path := visitedPath[curWord]
		//遍历当前单词的每个字符
		for i := 0; i < len(curWord); i++ {
			//从'a'到'z'遍历替换
			for k := 0; k < 26; k++ {
				//替换第i个字符
				newWord := curWord[:i] + string('a'+k) + curWord[i+1:]
				//新的字符串和终止字符串相同，返回当前遍历长度+1
				if strings.Compare(newWord, endWord) == 0 {
					return path + 1
				}
				//新单词在wordList中，但是没有访问过
				if _, ok := visitedPath[newWord]; !ok && wordListMap[newWord] == true {
					visitedPath[newWord] = path + 1
					Q = append(Q, newWord)
				}
			}
		}
	}
	return 0
}
