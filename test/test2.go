package main

import (
	"fmt"
	"strconv"
)

var nodeNum = 5

var visited = make(map[string]int)

//var node = []string{"S1", "S6", "S5", "S7"}
//
//var nodeMap = [][]int{
//	{-1, 161, -1, -1},
//	{171, -1, 181, -1},
//	{-1, 141, -1, 131},
//	{41, -1, -1, -1},
//}

var node = []string{"A", "B", "C", "D", "E"}

var nodeMap = [][]int{
	{-1, 8, 1, 2, -1},
	{-1, -1, -1, 3, 12},
	{-1, -1, -1, 6, 4},
	{7, 11, 10, -1, 5},
	{-1, 9, -1, -1, -1},
}

func findPath(start, end int, actionPath string) []string {
	res := make([]string, 0)
	if _, ok := visited[node[start]]; ok {
		return res
	}
	for i := 0; i < nodeNum; i++ {
		if nodeMap[start][i] == -1 {
			continue
		}
		if i == end {
			res = append(res, actionPath+strconv.Itoa(nodeMap[start][i])+" |")
			return res
		}
		visited[node[start]] = 1
		r := findPath(i, end, actionPath+strconv.Itoa(nodeMap[start][i])+"-->")
		delete(visited, node[start])
		if len(r) != 0 {
			res = append(res, r...)
		}
	}
	return res
}

func main() {
	fmt.Println(findPath(0, 4, ""))
}
