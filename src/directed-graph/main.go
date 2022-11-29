package main

const Inf = 999

type Matrix [][]int

type DistanceMapItem struct {
	distance     int
	previousNode int
	fixed        bool
}

type DistanceMap []DistanceMapItem

func getNearestUnfixedNode(m DistanceMap) int {
	tmp := DistanceMapItem{Inf, -1, false}
	res := -1
	for i, v := range m {
		if v.fixed {
			continue
		}
		if v.distance < tmp.distance {
			tmp = v
			res = i
		}
	}
	return res
}

func AllFixed(m DistanceMap) bool {
	for _, v := range m {
		if !v.fixed {
			return false
		}
	}
	return true
}

func DijkstraShortestPath(m Matrix, start int, dest int) (DistanceMap, []int) {
	route := []int{}
	currentNode := start
	distanceMap := make(DistanceMap, len(m))
	// 初期化
	for i := range distanceMap {
		distanceMap[i] = DistanceMapItem{Inf, -1, false}
	}

	for i := range m {
		if m[currentNode][i] == Inf {
			continue
		}
		distanceMap[i] = DistanceMapItem{m[currentNode][i], currentNode, false}
	}
	distanceMap[start].fixed = true
	currentNode = getNearestUnfixedNode(distanceMap)
	for {
		for i, v := range distanceMap {
			if i == currentNode {
				continue
			}
			if v.fixed {
				continue
			}
			pathLengthViaCurrentNode := distanceMap[currentNode].distance + m[currentNode][i]
			if pathLengthViaCurrentNode < v.distance {

				distanceMap[i] = DistanceMapItem{pathLengthViaCurrentNode, currentNode, false}
			}
		}
		route = append(route, currentNode)
		distanceMap[currentNode].fixed = true
		currentNode = getNearestUnfixedNode(distanceMap)
		if AllFixed(distanceMap) {
			break
		}
	}

	return distanceMap, route
}

func getChildrenFromMatrix(m Matrix, n int) []int {
	ret := []int{}
	for i, v := range m[n] {
		if v == 1 {
			ret = append(ret, i)
		}
	}
	return ret
}

func dfs(n int, m Matrix, sorted *[]int, visited *[]bool) {
	if (*visited)[n] {
		return
	}
	// visited[n] == false
	(*visited)[n] = true
	for _, v := range getChildrenFromMatrix(m, n) {

		dfs(v, m, sorted, visited)
	}
	*sorted = append(*sorted, n)
}

func ReverseSlice(s []int) []int {
	len := len(s)
	res := make([]int, len)
	for i, v := range s {
		res[len-1-i] = v
	}
	return res
}

func TopologicalSort(m Matrix) []int {
	sorted := []int{}
	visited := make([]bool, len(m))
	for i := 0; i < len(m); i++ {
		dfs(i, m, &sorted, &visited)
	}
	return ReverseSlice(sorted)
}
