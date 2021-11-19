package dst

import (
	"fmt"
)

// 图：一个图就是一些顶点的集合，这些顶点通过一系列边结对（连接）
// type LinkList struct{
// 	vertices []interface{}
// }
// 可使用邻接表或者矩阵来表达。
type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// add a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Printf("顶点 key %v 已经存在了, 不能再次添加\n", k)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// add edge
func (g *Graph) AddEdge(from, to int) {
	fVertex := g.getVertex(from)
	tVertex := g.getVertex(to)
	if fVertex == nil || tVertex == nil {
		fmt.Printf("起点 %v 或终点 %v 不存在\n", from, to)
		return
	}
	if contains(fVertex.adjacent, to) || contains(tVertex.adjacent, from) {
		fmt.Printf("此 (%v -> %v) Edge 已经存在了 \n", from, to)
		return
	}
	fVertex.adjacent = append(fVertex.adjacent, tVertex)
	tVertex.adjacent = append(tVertex.adjacent, fVertex)
}

// get vertex
func (g *Graph) getVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, key int) bool {
	for _, v := range s {
		if key == v.key {
			return true
		}
	}
	return false
}

// print
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %v: \n", v.key)
		for i, a := range v.adjacent {
			fmt.Println("第", i, "个顶点是：", a.key)
		}
	}
}
