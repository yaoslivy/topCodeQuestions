package main

//并差集的功能：
//1 寻找根节点，find(u int) 判断这个节点的祖先节点是哪个
//2 将两个节点接入到同一个集合，join(u, v int) 将两个节点连在同一个根节点上
//3 判断两个节点是否在同一个集合，same(u, v int) 判断两个节点是否在同一个根节点

//全局变量
var (
	n      = 1500 //节点数量
	father = make([]int, n)
)

//初始化并查集
func initialize() {
	for i := 0; i < n; i++ {
		father[i] = i
	}
}

//并查集里寻根
func find(u int) int {
	if u == father[u] {
		return u
	}
	father[u] = find(father[u])
	return father[u]
}

//将u-v这条边加入并查集
func join(u, v int) {
	//找到两个节点祖先节点，更改v的祖先节点为v的祖先节点，这样就连起了一条有向边
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u //v的上一个节点就是u
}

//判断u和v是否找到同一个根
func same(u, v int) bool {
	u = find(u)
	v = find(v)
	return u == v
}

func findRedundantConnection(edges [][]int) []int {
	//并查集
	initialize()
	for i := 0; i < len(edges); i++ {
		if same(edges[i][0], edges[i][1]) { //起始点和终止点的父节点是否相同
			return edges[i]
		} else {
			join(edges[i][0], edges[i][1])
		}
	}
	return []int{}
}
