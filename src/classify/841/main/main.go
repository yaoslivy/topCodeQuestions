package main

func canVisitAllRooms(rooms [][]int) bool {
	// return canVisitAllRoomsByDFS(rooms)
	return canVisitAllRoomsByBFS(rooms)
}

//广搜策略
func canVisitAllRoomsByBFS(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	//节点访问数组
	visited := make([]bool, len(rooms))
	//使用队列保存访问过的节点
	Q := make([]int, 0)
	Q = append(Q, 0)
	visited[0] = true
	for len(Q) != 0 {
		curNode := Q[0]
		Q = Q[1:]
		for _, val := range rooms[curNode] {
			if visited[val] == false {
				Q = append(Q, val)
				visited[val] = true
			}
		}
	}
	for _, val := range visited {
		if val == false {
			return false
		}
	}
	return true
}

//深搜策略
func canVisitAllRoomsByDFS(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	dfs(rooms, 0, visited)
	//检查深搜一遍后是否都访问到了
	for _, val := range visited {
		if val == false {
			return false
		}
	}
	return true
}
func dfs(rooms [][]int, i int, visited []bool) {
	if visited[i] {
		return
	}
	visited[i] = true
	keys := rooms[i]
	for _, val := range keys {
		dfs(rooms, val, visited)
	}
}
