package main

func numIslands(grid [][]byte) int {
	// return find(grid)
	return findByBFS(grid)
}

// bfs递归思路
func findByBFS(grid [][]byte) int {
	// 使用一个队列存储访问过的节点，再基于陆地广搜，将1-》0
	Q := make([]Node, 0)
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				grid[i][j] = '0'
				Q = append(Q, Node{i, j})
				for len(Q) != 0 {
					cur := Q[0]
					Q = Q[1:]
					if cur.i-1 >= 0 && grid[cur.i-1][cur.j] == '1' {
						Q = append(Q, Node{cur.i - 1, cur.j})
						grid[cur.i-1][cur.j] = '0'
					}
					if cur.j-1 >= 0 && grid[cur.i][cur.j-1] == '1' {
						Q = append(Q, Node{cur.i, cur.j - 1})
						grid[cur.i][cur.j-1] = '0'
					}
					if cur.i+1 < len(grid) && grid[cur.i+1][cur.j] == '1' {
						Q = append(Q, Node{cur.i + 1, cur.j})
						grid[cur.i+1][cur.j] = '0'
					}
					if cur.j+1 < len(grid[0]) && grid[cur.i][cur.j+1] == '1' {
						Q = append(Q, Node{cur.i, cur.j + 1})
						grid[cur.i][cur.j+1] = '0'
					}
				}
			}
		}
	}
	return res
}

type Node struct {
	i, j int
}

// dfs递归思路
func find(grid [][]byte) int {
	// 将为1的陆地通过dfs将所有相连的1-》0
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
}
