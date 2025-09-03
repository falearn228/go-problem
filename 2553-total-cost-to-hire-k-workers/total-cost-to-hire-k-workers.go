func totalCost(costs []int, k int, candidates int) int64 {
    h1 := &MinHeap{}
    h2 := &MinHeap{}
    heap.Init(h1)
    heap.Init(h2)
    var totalCost int64
    left := 0
    right := len(costs) - 1
    for range k {
        for h1.Len() < candidates && left <= right {
            heap.Push(h1, [2]int{costs[left], left})
            left++
        }
        for h2.Len() < candidates && left <= right {
            heap.Push(h2, [2]int{costs[right], right})
            right--
        }

        lCost, rCost := -1, -1
        if h1.Len() > 0 {
            lCost = (*h1)[0][0]
        }
        if h2.Len() > 0 {
            rCost = (*h2)[0][0]
        }

        
        if lCost != -1 && (rCost == -1 || lCost <= rCost) {
            costAndInd := heap.Pop(h1).([2]int)
            totalCost += int64(costAndInd[0])
        } else if rCost != -1 {
            costAndInd := heap.Pop(h2).([2]int)
            totalCost += int64(costAndInd[0])
        }
    }

    return totalCost
}

type MinHeap [][2]int

func (m MinHeap) Len() int { return len(m)}
func (m MinHeap) Less(i, j int) bool {
    if m[i][0] != m[j][0] {
        return m[i][0] < m[j][0]
    }
    return m[i][1] < m[j][1]
}
func (m MinHeap) Swap(i, j int) {  m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(value any) {
    *m = append(*m, value.([2]int))
}

func (m *MinHeap) Pop() any {
    old := *m
    n := len(old)
    top := old[n-1]
    *m = old[:n-1]
    return top
}

