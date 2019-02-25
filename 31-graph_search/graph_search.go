package main

import (
	"container/list"
	"fmt"
)

// Graph 图结构
type Graph struct {
	adj []*list.List
	v   int
}

func main() {
	fmt.Println("hello world.")
	fmt.Println("init graph...")
	gf := NewGraph(9)
	gf.addEdge(0, 1)
	gf.addEdge(0, 3)
	gf.addEdge(1, 2)
	gf.addEdge(1, 4)
	gf.addEdge(3, 6)
	gf.addEdge(4, 5)
	gf.addEdge(6, 7)
	gf.addEdge(5, 8)
	/** 图结构，无向图
	0--1--2
	|  |  |
	3  4--5
	|     |
	6--7  8
	*/

	fmt.Println("init Done.\nStart Search")
	fmt.Println("0 -> 7")
	gf.BFS(0, 7)
	fmt.Println("1 -> 3")
	gf.BFS(1, 3)
	fmt.Println("7 -> 8")
	gf.BFS(7, 8)
}

// NewGraph 初始化Graph
func NewGraph(n int) *Graph {
	gf := &Graph{}
	gf.v = n
	gf.adj = make([]*list.List, n)
	for i := range gf.adj {
		gf.adj[i] = list.New()
	}
	return gf
}
func (gf *Graph) addEdge(from, to int) {
	gf.adj[from].PushBack(to)
	gf.adj[to].PushBack(from)
}

// BFS 广度优先查找
// 使用BFS查找是否存在 from 到 to 的边
func (gf *Graph) BFS(from, to int) {
	if from == to {
		fmt.Println("no need to find the same point")
		return
	}
	fmt.Println("启用BFS 广度优先查找")
	var queue []int
	visited := make([]bool, gf.v)
	prev := make([]int, gf.v)
	isFound := false

	queue = append(queue, from)
	for v := range prev {
		prev[v] = -1
	}

	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		visited[top] = true
		for e := gf.adj[top].Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			// fmt.Println("print e.value", k)
			if !visited[k] {
				queue = append(queue, k)
				prev[k] = top
				if k == to {
					isFound = true
					break
				}
			}
		}
	}
	// 打印结果
	if isFound {
		fmt.Println("found...")
		link := fmt.Sprintf("%d", to)
		for v := prev[to]; v != -1; v = prev[v] {
			link = fmt.Sprintf("%d->%s", v, link)
		}
		fmt.Println(link)
	} else {
		fmt.Println("not found")
	}
}
