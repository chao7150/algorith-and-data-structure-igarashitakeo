package main

const Inf = 999

type Matrix [][]int

func Fill(len int, val int) []int {
	res := make([]int, len)
	for i := range res {
		res[i] = val
	}
	return res
}

func CreateSparseAdjacencyMatrix(len int) Matrix {
	res := make([][]int, len)
	for i := range res {
		res[i] = Fill(len, Inf)
	}
	for i := range res {
		res[i][i] = 0
	}
	return res
}

func SelectAddingNode(remainingNodes []int, costTable []int) int {
	node := -1
	cost := Inf
	for _, v := range remainingNodes {
		if costTable[v] < cost {
			node = v
			cost = costTable[v]
		}
	}
	return node
}

func RemoveFirst(s []int, r int) []int {
	idx := -1
	for i, v := range s {
		if v == r {
			idx = i
			break
		}
	}
	return append(s[:idx], s[idx+1:]...)
}

func Prim(m Matrix) Matrix {
	start := 0
	costTable := m[start]
	nearestNodesInTree := Fill(len(m), start)
	remainingNodes := []int{}
	// start以外のノードを入れる
	for i := range m {
		if i != start {
			remainingNodes = append(remainingNodes, i)
		}
	}
	minSpanningTree := CreateSparseAdjacencyMatrix(len(m))
	for 0 < len(remainingNodes) {
		addingNode := SelectAddingNode(remainingNodes, costTable)
		nearestNodeInTree := nearestNodesInTree[addingNode]
		remainingNodes = RemoveFirst(remainingNodes, addingNode)
		// このループで追加されるエッジのコスト
		cost := m[addingNode][nearestNodeInTree]
		minSpanningTree[addingNode][nearestNodeInTree] = cost
		minSpanningTree[nearestNodeInTree][addingNode] = cost
		for _, v := range remainingNodes {
			if costTable[v] > m[v][addingNode] {
				costTable[v] = m[v][addingNode]
				nearestNodesInTree[v] = addingNode
			}
		}
	}
	return minSpanningTree
}
