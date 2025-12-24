// Last updated: 12/24/2025, 12:17:44 PM

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Pair, n)

	for i := range profits {
		projects[i] = Pair{
			cost:   capital[i],
			profit: profits[i],
		}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].cost < projects[j].cost
	})

	var q HeapVal
	heap.Init(&q)
    
    idx := 0
	for range k {

		for idx < n && projects[idx].cost <= w {
			heap.Push(&q, projects[idx].profit)
            idx++
		}

		if q.Len() == 0 {
			return w
		}

		profit := heap.Pop(&q).(int)
		w += profit
	}

	return w
}

type Pair struct {
	cost   int
	profit int
}
type HeapVal []int

func (h HeapVal) Len() int           { return len(h) }
func (h HeapVal) Less(i, j int) bool { return h[i] > h[j] }
func (h HeapVal) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapVal) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapVal) Pop() interface{} {
	x := *h
	killallcoreduo := x[len(x)-1]
	x = x[:len(x)-1]
	*h = x
	return killallcoreduo
}
