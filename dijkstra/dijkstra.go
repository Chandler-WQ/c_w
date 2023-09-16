package dijkstra

import (
	"math"
)

type Graph struct {
	Distance [][]float64
}

func (graph *Graph) AddEdge(from, to int, distance float64) {
	graph.Distance[from][to] = distance
	graph.Distance[to][from] = distance
}

func NewGraph(num int) Graph {
	distancess := make([][]float64, 0, num)
	for i := 0; i < num; i++ {
		distances := make([]float64, num)
		for j := 0; j < num; j++ {
			distances[j] = math.Inf(1)
			if i == j {
				distances[j] = 0
			}
		}
		distancess = append(distancess, distances)
	}
	return Graph{
		Distance: distancess,
	}
}

func (g *Graph) Length(i, j int) float64 {
	return g.Distance[i][j]
}

// 找寻点到一个线段之间的最短路径
func (g *Graph) ShortestPathInPointPath(source int, targets [2]int) ([]int, float64) {
	min := math.Inf(1)
	path := []int{}
	for _, target := range targets {
		p, dist := g.ShortestPathInPoints(source, target)
		if dist < min {
			min = dist
			path = p
		}
	}
	return path, min
}

// 找寻两个线段之间的最短距离
func (g *Graph) ShortestPathInPaths(sources, targets [2]int) ([]int, float64) {
	min := math.Inf(1)
	path := []int{}
	for _, source := range sources {
		for _, target := range targets {
			p, dist := g.ShortestPathInPoints(source, target)
			if dist < min {
				min = dist
				path = p
			}
		}
	}
	return path, min
}

// 找寻两个点之间的最短距离
func (g *Graph) ShortestPathInPoints(source, target int) ([]int, float64) {
	numNodes := len(g.Distance)
	visited := make([]bool, numNodes)
	Distances := make([]float64, numNodes)
	previous := make([]int, numNodes)

	for i := 0; i < numNodes; i++ {
		Distances[i] = math.Inf(1)
		previous[i] = -1
	}

	Distances[source] = 0

	for count := 0; count < numNodes-1; count++ {
		u := findMinDistance(Distances, visited)
		visited[u] = true

		for v := 0; v < numNodes; v++ {
			if !visited[v] && g.Distance[u][v] >= 0 && Distances[u]+g.Distance[u][v] < Distances[v] {
				Distances[v] = Distances[u] + g.Distance[u][v]
				previous[v] = u
			}
		}
	}

	path := buildPath(previous, target)
	length := Distances[target]

	return path, length
}

func findMinDistance(Distances []float64, visited []bool) int {
	min := math.Inf(1)
	minIndex := -1

	for i, dist := range Distances {
		if !visited[i] && dist < min {
			min = dist
			minIndex = i
		}
	}

	return minIndex
}

func buildPath(previous []int, target int) []int {
	path := []int{target}
	for previous[target] != -1 {
		target = previous[target]
		path = append([]int{target}, path...)
	}
	return path
}
