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
	gf.addEdge(2, 5)
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
	fmt.Println("0 -> 8")
	gf.BFS(0, 8)
	fmt.Println("1 -> 3")
	gf.BFS(1, 3)
	fmt.Println("7 -> 8")
	gf.BFS(7, 8)

	fmt.Println("0 -> 7")
	gf.DFS(0, 7)
	fmt.Println("0 -> 8")
	gf.DFS(0, 8)
	fmt.Println("0 -> 3")
	gf.DFS(0, 3)
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
	fmt.Println("====start BFS=====")
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

// 用于DFS中标记是否已找到
var dFound bool

// DFS 深度优先查找
func (gf *Graph) DFS(from, to int) {
	fmt.Println("====start DFS=====")
	dFound = false
	var queue []int
	visited := make([]bool, gf.v)
	prev := make([]int, gf.v)
	queue = append(queue, from)
	visited[0] = true
	for k := range prev {
		prev[k] = -1
	}
	if from == to {
		fmt.Println("the same node does not need to find.")
		return
	}
	gf.recurDfs(from, to, queue, visited, prev)
	if dFound {
		fmt.Println("found...")
		link := fmt.Sprintf("%d", to)
		for v := prev[to]; v != -1; v = prev[v] {
			link = fmt.Sprintf("%d->%s", v, link)
		}
		fmt.Println(link)
	} else {
		fmt.Println("not found.")
	}
}

// recurDfs 递归查找
func (gf *Graph) recurDfs(from, to int, queue []int, visited []bool, prev []int) {
	if dFound || len(queue) < 1 {
		return
	}
	qlen := len(queue)
	topq := queue[qlen-1 : qlen]
	top := topq[0]
	queue = queue[1:]
	for el := gf.adj[top].Front(); el != nil; el = el.Next() {
		val := el.Value.(int)
		if !visited[val] {
			prev[val] = top
			visited[val] = true
			queue = append(queue, val)
			if val == to {
				dFound = true
				break
			}
		}
		gf.recurDfs(from, to, queue, visited, prev)
	}
}
