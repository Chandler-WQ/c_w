package cw

import (
	"fmt"
	"sort"
)

type Graph struct {
	Distance [][]float64
}

type Save struct {
	i       int
	j       int
	saveIng float64
}

func (graph *Graph) String() string {
	return fmt.Sprintf("Distance: %v", graph.Distance)
}

func (graph *Graph) AddEdge(from, to int, distance float64) {
	graph.Distance[from][to] = distance
	graph.Distance[to][from] = distance
}

func NewGraph(distance [][]float64) *Graph {
	return &Graph{
		Distance: distance,
	}
}

type Route struct {
	Nodes []int
}

func (r *Route) Length(graph *Graph) float64 {
	length := 0.0
	for i := 0; i < len(r.Nodes)-1; i++ {
		from := r.Nodes[i]
		to := r.Nodes[i+1]
		length += graph.Distance[from][to]
	}
	return length
}

func (r *Route) InsertNode(index, node int) {
	r.Nodes = append(r.Nodes[:index], append([]int{node}, r.Nodes[index:]...)...)
}

func (r *Route) RemoveNode(index int) int {
	node := r.Nodes[index]
	r.Nodes = append(r.Nodes[:index], r.Nodes[index+1:]...)
	return node
}

// func (r *Route) MergeWith(other *Route, index1 int) {
// 	r.Nodes = append(r.Nodes[:index1], append(other.Nodes, r.Nodes[0:]...)...)
// }

func (r *Route) Optimize(graph *Graph) {
	improvement := true
	for improvement {
		improvement = false
		for i := 1; i < len(r.Nodes)-2; i++ {
			for j := i + 1; j < len(r.Nodes)-1; j++ {
				if r.SwapImproves(i, j, graph) {
					r.SwapNodes(i, j)
					improvement = true
				}
			}
		}
	}
}

func (r *Route) SwapImproves(i, j int, graph *Graph) bool {
	node1 := r.Nodes[i]
	node2 := r.Nodes[i+1]
	node3 := r.Nodes[j]
	node4 := r.Nodes[j+1]

	oldDistance := graph.Distance[node1][node2] + graph.Distance[node3][node4]
	newDistance := graph.Distance[node1][node3] + graph.Distance[node2][node4]

	return newDistance < oldDistance
}

func (r *Route) SwapNodes(i, j int) {
	r.Nodes[i+1], r.Nodes[j] = r.Nodes[j], r.Nodes[i+1]
}

func (r *Route) Copy() *Route {
	newRoute := &Route{
		Nodes: make([]int, len(r.Nodes)),
	}
	copy(newRoute.Nodes, r.Nodes)
	return newRoute
}

func (r *Route) String() string {
	return fmt.Sprintf("%v", r.Nodes)
}

func (graph *Graph) FindOptimalRoute() []*Route {
	numNodes := len(graph.Distance)

	routes := make([]*Route, numNodes-1)
	for i := 0; i < numNodes-1; i++ {
		routes[i] = &Route{
			Nodes: []int{i + 1},
		}
	}
	saves := make([]Save, 0, numNodes-1)
	for i := 0; i < len(routes)-1; i++ {
		for j := i + 1; j < len(routes); j++ {
			savings := graph.CalculateSavings(routes[i], routes[j])
			if savings > 0 {
				saves = append(saves, Save{i + 1, j + 1, savings})
			}
		}
	}
	sort.Slice(saves, func(i, j int) bool {
		return saves[i].saveIng > saves[j].saveIng
	})
	for _, save := range saves {
		if len(routes) == 1 {
			break
		}
		var startRouteIndex, endRouteIndex int = -1, -1
		for i, route := range routes {
			if save.i == route.Nodes[0] {
				startRouteIndex = i
			}
			if save.j == route.Nodes[len(route.Nodes)-1] {
				endRouteIndex = i
			}
			if startRouteIndex >= 0 && endRouteIndex >= 0 {
				routes[startRouteIndex].Nodes = append(routes[startRouteIndex].Nodes, routes[endRouteIndex].Nodes...)
				routes = append(routes[:endRouteIndex], routes[endRouteIndex+1:]...)
				break
			}
		}
	}
	for i := range routes {
		routes[i].Nodes = append([]int{0}, routes[i].Nodes...)
		routes[i].Nodes = append(routes[i].Nodes, 0)
	}
	return routes
}

func (graph *Graph) CalculateSavings(route1, route2 *Route) float64 {
	savings := 0.0
	for _, node1 := range route1.Nodes {
		for _, node2 := range route2.Nodes {
			savings += graph.Distance[node1][0] + graph.Distance[0][node2] - graph.Distance[node1][node2]
		}
	}
	return savings
}
