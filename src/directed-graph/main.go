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

func DijkstraShortestPath(m Matrix, start int) DistanceMap {
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
		distanceMap[currentNode].fixed = true
		currentNode = getNearestUnfixedNode(distanceMap)
		if AllFixed(distanceMap) {
			break
		}
	}

	return distanceMap
}
