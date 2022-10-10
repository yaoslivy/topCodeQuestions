package main

//并查集
var (
	n      = 1005
	father = make([]int, n)
)

//初始化并查集
func initialize() {
	for i := 0; i < len(father); i++ {
		father[i] = i
	}
}

//查找节点的父节点
func find(u int) int {
	if u == father[u] {
		return u
	}
	father[u] = find(father[u])
	return father[u]
}

//将边u->v加入集合
func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u
}

//判断两个节点的父节点是否相同
func same(u, v int) bool {
	if find(u) == find(v) {
		return true
	}
	return false
}

//isTreeAfterRemoveEdge 删除一条边后判断是不是树
func isTreeAfterRemoveEdge(edges [][]int, deleteEdge int) bool {
	//初始化并查集
	initialize()
	for i := 0; i < len(edges); i++ {
		if i == deleteEdge { //删除的边的起始点在集合中的下标
			continue
		}
		if same(edges[i][0], edges[i][1]) { //构成有向环，一定不是树
			return false
		}
		join(edges[i][0], edges[i][1])
	}
	return true
}

// getRemoveEdge 在有向图中找到删除的那条边，使其成为树
func getRemoveEdge(edges [][]int) []int {
	//初始化并查集
	initialize()
	for i := 0; i < len(edges); i++ {
		if same(edges[i][0], edges[i][1]) { //构成了有向环，即是要删除的边
			return edges[i]
		}
		join(edges[i][0], edges[i][1])
	}
	return []int{}
}

func findRedundantDirectedConnection(edges [][]int) []int {
	//先统计各个节点的入度
	inDegree := make([]int, len(father))
	for i := 0; i < len(edges); i++ {
		inDegree[edges[i][1]] += 1
	}
	//记录入度为2的节点对对应边的下标，倒过来找，则后的边先加入集合
	twoDegree := make([]int, 0)
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			twoDegree = append(twoDegree, i)
		}
	}
	// 如果有入度为2的节点，一定是在两条边里面删除一个，看删除哪个可以构成树
	if len(twoDegree) > 0 {
		if isTreeAfterRemoveEdge(edges, twoDegree[0]) {
			return edges[twoDegree[0]]
		}
		return edges[twoDegree[1]]
	}
	// 没有入度为2的情况，一定是有向环，找到构成环的边返回即可
	return getRemoveEdge(edges)
}
