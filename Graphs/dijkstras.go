// Dijkstra's algorithm
/*
	Dijkstra's algorithm is a pathfinding algorithm that finds the shortest path between two points in a graph.
	Here's how it works:

	1. Start at the starting point (let's call it "A") and mark its distance as 0.

	2. Visit each of its neighbors (the nodes it's connected to) and mark their distances as their weight
	   (the distance between node A and its neighbor).

	3. Choose the neighbor with the shortest distance and mark it as "visited."

	4. Visit each of its neighbors that haven't been visited yet and calculate their distances by adding the weight
	   of the edge between them and the current node to the current node's distance. If this distance is shorter than the current distance marked for the node, replace it with the shorter distance.

	5. Repeat steps 3 and 4 until you reach the destination node or until all nodes have been visited.

	6. Once you've reached the destination node, the shortest path can be determined by tracing back through the
	   visited nodes from the destination to the starting point, following the path of the shortest distances.

	Time complexity:

		In the worst case, Dijkstra's algorithm visits every vertex and every edge in the graph.
		Each vertex is visited once, and each edge is relaxed once.
		Therefore, the time complexity of Dijkstra's algorithm is O(V + E * log V), where V is the number of vertices and E is the number of edges in the graph.
		Note that this assumes a priority queue is used to efficiently extract the vertex with the smallest distance. If an array is used instead, the time complexity would be O(V^2 + E).

	Space complexity:

		Dijkstra's algorithm uses two data structures: a priority queue and an array to keep track of the shortest distances from the source vertex to all other vertices.
		The priority queue stores at most V vertices at a time.
		The array has V entries.
		Therefore, the space complexity of Dijkstra's algorithm is O(V).
*/
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// priorityQueue is a heap-based priority queue of items
type priorityQueue []*item

// item represents an item in the priority queue
type item struct {
    value    string  // value of the item
    priority int     // priority of the item (i.e., its distance from the start vertex)
}

// Len returns the number of items in the priority queue
func (pq priorityQueue) Len() int {
    return len(pq)
}

// Less returns true if item i has higher priority than item j
func (pq priorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

// Swap swaps the positions of items i and j in the priority queue
func (pq priorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an item to the priority queue
func (pq *priorityQueue) Push(x interface{}) {
    item := x.(*item)
    *pq = append(*pq, item)
}

// Pop removes the item with the highest priority from the priority queue and returns it
func (pq *priorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

// Dijkstra's algorithm finds the shortest path between a start vertex and all other vertices in a weighted graph.
// It returns a map of the shortest distance to each vertex and the previous vertex on the shortest path.
// If a vertex is unreachable, its distance is set to infinity.
func Dijkstra(graph map[string]map[string]int, start string) (map[string]int, map[string]string) {
    dist := make(map[string]int)      // distances from start vertex to each vertex
    prev := make(map[string]string)   // previous vertex on the shortest path from start to each vertex
    queue := make(priorityQueue, 0)   // priority queue to track next vertex to visit
    
    // initialize distances and previous vertices
    for v := range graph {
        dist[v] = math.MaxInt32    // initialize all distances to infinity
        prev[v] = ""               // initialize all previous vertices to empty
    }
    dist[start] = 0                // distance from start vertex to itself is 0
    
    // add start vertex to priority queue
    heap.Push(&queue, &item{value: start, priority: 0})
    
    // loop until priority queue is empty
    for queue.Len() > 0 {
        // remove vertex with minimum distance from priority queue
        u := heap.Pop(&queue).(*item).value
        
        // loop through neighbors of current vertex
        for v, w := range graph[u] {
            alt := dist[u] + w   // calculate alternate distance to neighbor
            
            // if alternate distance is shorter than current distance, update distances and previous vertices
            if alt < dist[v] {
                dist[v] = alt
                prev[v] = u
                
                // add neighbor to priority queue
                heap.Push(&queue, &item{value: v, priority: alt})
            }
        }
    }
    
    return dist, prev
}

func main() {

	graph := map[string]map[string]int{
		"A": {"B": 2, "C": 3},
		"B": {"C": 1, "D": 1},
		"C": {"D": 4},
		"D": {"C": 2},
	}
	start := "A"
	distances, _ := Dijkstra(graph, start)
	fmt.Println("Shortest distances from", start, "to all other nodes:")
	for node, distance := range distances {
    fmt.Println(node, distance)
}
}