// Last updated: 11/7/2025, 2:47:18 PM
package main

import (
	"container/heap"
	"math"
)

type pair struct {
	node  int
	price int
	stops int // max_value = k
}
type pqueue []pair

func (p pqueue) Len() int           { return len(p) }
func (p pqueue) Less(i, j int) bool { return p[i].price < p[j].price }
func (p pqueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pqueue) Pop() any {
	old := *p
	val := old[len(old)-1]
	old = old[:len(old)-1]
	*p = old
	return val
}
func (p *pqueue) Push(val any) {
	*p = append(*p, val.(pair))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj := make(map[int][]pair, len(flights))

	for i := range flights {
		from, to, price := flights[i][0], flights[i][1], flights[i][2]
		adj[from] = append(adj[from], pair{node: to, price: price})

	}

	stopsMap := make(map[int]int, len(flights))
	for i := 0; i < n; i++ {
		stopsMap[i] = math.MaxInt
	}

	pq := &pqueue{pair{src, 0, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(pair)

		if curr.node == dst {
			return curr.price
		}
		if curr.stops > k {
			continue
		}

		if curr.stops >= stopsMap[curr.node] {
			continue
		}
		stopsMap[curr.node] = curr.stops
		for _, neighbors := range adj[curr.node] {
			// neighbors.price += curr.price
			// neighbors.stops++
			// heap.Push(pq, neighbors)
			heap.Push(pq, pair{
				node:  neighbors.node,
				price: curr.price + neighbors.price, // Новая цена = текущая + цена перелета
				stops: curr.stops + 1,               // Увеличиваем кол-во остановок на 1
			})
		}
	}

	return -1
}
