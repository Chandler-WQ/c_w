package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Chandler-WQ/c_w/cw"
	"github.com/Chandler-WQ/c_w/dijkstra"
)

type Edge struct {
	Name     string
	From     int
	To       int
	Distance float64
}

func main() {
	// 构造初始的铁路网
	originGraph := dijkstra.NewGraph(13)
	originGraph.AddEdge(0, 1, 15)
	originGraph.AddEdge(0, 4, 29)
	originGraph.AddEdge(0, 5, 40)
	originGraph.AddEdge(8, 5, 15)
	originGraph.AddEdge(8, 7, 14)
	originGraph.AddEdge(8, 12, 29)
	originGraph.AddEdge(11, 12, 23)
	originGraph.AddEdge(7, 12, 34)
	originGraph.AddEdge(7, 4, 15)
	originGraph.AddEdge(3, 4, 7)
	originGraph.AddEdge(3, 6, 19)
	originGraph.AddEdge(3, 2, 21)
	originGraph.AddEdge(11, 10, 22)
	originGraph.AddEdge(1, 2, 24)
	originGraph.AddEdge(2, 6, 22)
	originGraph.AddEdge(10, 6, 19)
	originGraph.AddEdge(9, 6, 20)
	originGraph.AddEdge(9, 10, 18)
	originGraph.AddEdge(7, 10, 20)

	type Edge struct {
		Points [2]int
		Name   string
		Length float64
	}

	edge := []Edge{
		// {
		// 	Points: [2]int{1, 1},
		// 	Name:   "Depot_1",
		// 	Length: 0,
		// },
		{
			Points: [2]int{8, 8},
			Name:   "Depot_2",
			Length: 0,
		},
		// {
		// 	Points: [2]int{0, 1},
		// 	Name:   "A",
		// 	Length: 15,
		// },
		// {
		// 	Points: [2]int{2, 3},
		// 	Name:   "B",
		// 	Length: 21,
		// },
		{
			Points: [2]int{4, 7},
			Name:   "C",
			Length: 15,
		},
		{
			Points: [2]int{5, 8},
			Name:   "D",
			Length: 15,
		},
		{
			Points: [2]int{9, 10},
			Name:   "E",
			Length: 18,
		},
		{
			Points: [2]int{11, 12},
			Name:   "F",
			Length: 23,
		},
	}

	newGraph := dijkstra.NewGraph(len(edge))
	for i := 0; i < len(edge); i++ {
		for j := i; j < len(edge); j++ {
			if i == j {
				newGraph.AddEdge(i, j, 0)
				continue
			}
			_, length := originGraph.ShortestPathInPaths(edge[i].Points, edge[j].Points)
			// 维修段坍缩成点之间的距离为每个维修段之间的最短距离+两侧维修段长度
			length = length + (edge[i].Length + edge[j].Length)
			newGraph.AddEdge(i, j, length)
		}
	}
	// fmt.Println(newGraph.Distance)

	cwGraph := cw.NewGraph(newGraph.Distance)
	optimalRoutes := cwGraph.FindOptimalRoute()
	for _, optimalRoute := range optimalRoutes {
		edgeName := []string{}
		for i := 0; i < len(optimalRoute.Nodes); i++ {
			edgeName = append(edgeName, edge[optimalRoute.Nodes[i]].Name)
		}
		fmt.Println("维修顺序是:", strings.Join(edgeName, "->"))

		length := 0
		path := []string{}
		for i := 0; i < len(optimalRoute.Nodes)-1; i++ {
			fromEdge := edge[optimalRoute.Nodes[i]]
			toEdge := edge[optimalRoute.Nodes[i+1]]
			path = append(path, fromEdge.Name)
			seqPath, segLength := originGraph.ShortestPathInPaths(fromEdge.Points, toEdge.Points)
			seqPathStr := Int2StrSlice(seqPath)
			if len(path) > 0 && path[len(path)-1] == seqPathStr[0] {
				path = path[:len(path)-1]
			}
			path = append(path, seqPathStr...)
			path = append(path, toEdge.Name)
			length += int(segLength)
		}
		fmt.Printf("不考虑维修区段成本，途中长度为: %v km,路径为:%v\n", length, strings.Join(path, "->"))
	}

}

func Int2StrSlice(in []int) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = strconv.Itoa(v)
	}
	return out
}

//[[0 21.5 7.5 43] [21.5 0 29 36.5] [7.5 29 0 50.5] [43 36.5 50.5 0]],
