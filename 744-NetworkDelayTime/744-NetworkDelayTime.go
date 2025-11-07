// Last updated: 11/7/2025, 2:47:22 PM
type pair struct {
	node int
	w    int
}

type PrQueue []pair

func (h PrQueue) Len() int {
	return len(h)
}
func (h PrQueue) Less(i, j int) bool {
	return h[i].w < h[j].w
}
func (h PrQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *PrQueue) Push(val any) {
	*h = append(*h, val.(pair))
}
func (h *PrQueue) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func networkDelayTime(times [][]int, n int, k int) int {
	// u -> pair[v, w]
    minTime := make(map[int]int, len(times))
	dists := make(map[int][]pair, len(times))
    for _,t := range times {
        u, v, w := t[0], t[1], t[2]
        dists[u] = append(dists[u], pair{node: v, w: w})
    }
    pq := &PrQueue{pair{k, 0}}
    heap.Init(pq)
	maxAnswer := 0
    for pq.Len() > 0 {
        p := heap.Pop(pq).(pair)
        currNode, time := p.node, p.w
        
        if _, visited := minTime[currNode]; visited {
            continue
        }
        minTime[currNode] = time

        for _, edge := range dists[currNode] {
            neighbor, cost := edge.node, edge.w

            if _, ok := minTime[neighbor]; !ok {
                newTime := cost + time
                heap.Push(pq, pair{neighbor, newTime})
            }
        }
    }

    if len(minTime) != n {
        return -1
    }
    for _, cost := range minTime {
        maxAnswer = max(cost, maxAnswer)
    }

    return maxAnswer
}
