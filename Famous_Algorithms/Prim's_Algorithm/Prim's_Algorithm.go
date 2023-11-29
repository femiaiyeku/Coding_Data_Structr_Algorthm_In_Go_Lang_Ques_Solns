/*


You're given a list of edges representing a connected, weighted, undirected graph with at least one node.
The given list is what's called an adjacency list, and it represents a graph. The number of vertices in the graph is equal to
the length of edges, where each index in edges contains vertex's siblings, in no particular order. Each of
these siblings is an array of length two, with the first value denoting the index in the list that this vertex is connected to,
and and the second value denoting the weight of the edge. Note that this graph is undirected, meaning that if a vertex
appears in the edge list of another vertex, then the inverse will also be true (along with the same weight).
Write a function implementing Prim's Algorithm to return a new edges array that represents a minimum spanning
tree. A minimum spanning tree is a tree containing all of the vertices of the original graph and a subset of the edges.
These edges should connect all of the vertices with the minimum total edge weight and without generating any cycles.
Note that the graph represented by edges won't contain any self-loops (vertices that have an outbound edge to
themselves) and will only have positively weighted edges (l.e., no negative distances). The graph will always have at least
one node.
If you're unfamiliar with Prim's algorithm, we recommend watching the Conceptual Overview section of this question's
video explanation before starting to code.


Sample Input:
edges = [
    [[1, 7]],
    [[2, 8], [3, 3], [4, 6]],
    [[3, 4]],
    [[4, 2]],
    [],
    []
]

Sample Output:
[
    [[1, 7]],
    [[0, 7], [3, 3], [4, 6]],
    [[3, 4]],
    [[1, 3]],
    [[1, 6]],
    []
]











*/



package main

import (
    "container/heap"
    "fmt"
)

type Edge struct {
    from, to  int
    weight   int
}

type EdgeHeap []Edge

func (eh EdgeHeap) Len() int {
    return len(eh)
}

func (eh EdgeHeap) Less(i, j int) bool {
    return eh[i].weight < eh[j].weight
}

func (eh EdgeHeap) Swap(i, j int) {
    eh[i], eh[j] = eh[j], eh[i]
}

func (eh *EdgeHeap) Push(x interface{}) {
    *eh = append(*eh, x.(Edge))
}

func (eh *EdgeHeap) Pop() interface{} {
    n := len(*eh)
    edge := (*eh)[n-1]
    *eh = (*eh)[:n-1]
    return edge
}

func primMST(edges [][]Edge) [][]Edge {
    n := len(edges)
    visited := make([]bool, n)
    mst := make([][]Edge, n)

    for i := 0; i < n; i++ {
        if !visited[i] {
            mst[i] = []Edge{}
            primMSTHelper(edges, i, visited, mst)
        }
    }

    return mst
}

func primMSTHelper(edges [][]Edge, start int, visited []bool, mst [][]Edge) {
    visited[start] = true

    pq := EdgeHeap{}
    heap.Init(&pq)

    for _, edge := range edges[start] {
        if !visited[edge.to] {
            heap.Push(&pq, Edge{start, edge.to, edge.weight})
        }
    }

    for pq.Len() > 0 {
        edge := heap.Pop(&pq).(Edge)
        from, to := edge.from, edge.to
        weight := edge.weight

        if !visited[to] {
            visited[to] = true
            mst[to] = []Edge{}

            mst[from] = append(mst[from], Edge{from, to, weight})
            mst[to] = append(mst[to], Edge{to, from, weight}) // Undirected graph

            for _, edge := range edges[to] {
                if !visited[edge.to] {
                    heap.Push(&pq, Edge{to, edge.to, edge.weight})
                }
            }
        }
    }
}

func main() {
    edges := [][]Edge{
        {{1, 7}},
        {{2, 8}, {3, 3}, {4, 6}},
        {{3, 4}},
        {{4, 2}},
        [],
        [],
    }

    mst := primMST(edges)
    fmt.Println(mst)
}
