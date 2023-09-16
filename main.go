package main

import (
	"fmt"

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
	}

	edge := []Edge{
		{
			Points: [2]int{8, 8},
			Name:   "Depot_2",
		},
		{
			Points: [2]int{4, 7},
			Name:   "C",
		},
		{
			Points: [2]int{5, 8},
			Name:   "D",
		},
		{
			Points: [2]int{9, 10},
			Name:   "E",
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
			// 维修段坍缩成点之间的距离为每个维修段之间的最短距离+两侧维修段长度的一半
			length = length + (originGraph.Length(edge[i].Points[0], edge[i].Points[1])+originGraph.Length(edge[j].Points[0], edge[j].Points[1]))/2
			newGraph.AddEdge(i, j, length)
		}
	}

	cwGraph := cw.NewGraph(newGraph.Distance)
	optimalRoute := cwGraph.FindOptimalRoute()
	optimalRoute.Optimize(cwGraph)
	edgeName := []string{}
	for i := 0; i < len(optimalRoute.Nodes); i++ {
		edgeName = append(edgeName, edge[optimalRoute.Nodes[i]].Name)
	}
	fmt.Println("维修顺序是:", edgeName)

	// length := 0
	// path := []int{}
	// for i := 0; i < len(optimalRoute.Nodes)-2; i++ {
	// 	fromEdge := edge[i]
	// 	toEdge := edge[i+1]
	// 	seqPath, segLength := originGraph.ShortestPathInPaths(fromEdge.Points, toEdge.Points)
	// 	fmt.Println("seqPath,segLength", seqPath, segLength)
	// 	path = append(path, seqPath...)
	// 	length += int(segLength)
	// }
	// fmt.Println("length,path", length, path)
}

//[[0 21.5 7.5 43] [21.5 0 29 36.5] [7.5 29 0 50.5] [43 36.5 50.5 0]],
