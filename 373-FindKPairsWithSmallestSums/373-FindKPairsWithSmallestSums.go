// Last updated: 12/24/2025, 12:17:55 PM
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    var h HeapVal
    heap.Init(&h)
    res := make([][]int, k)

    limit := len(nums1)
	if k < limit {
		limit = k
	}

    for i := range limit {
        heap.Push(&h, Pair{
            i1: i,
            i2: 0,
            sum: nums1[i] + nums2[0] ,
        })
    }

    m := len(nums2)
    idxRes := 0
    for range k {
        if h.Len() == 0 {
            break
        }
        curr := heap.Pop(&h).(Pair)

        res[idxRes] = []int{nums1[curr.i1], nums2[curr.i2]}
        idxRes++

        nextIdx2 := curr.i2 + 1
        if nextIdx2 < m {
            currSum := nums1[curr.i1] + nums2[nextIdx2]
            heap.Push(&h, Pair{sum: currSum, i1: curr.i1, i2: nextIdx2})
        }
    }

    return res
}
type Pair struct {
    i1 int
    i2 int
    sum int
}
type HeapVal []Pair

func (h HeapVal) Len() int           { return len(h) }
func (h HeapVal) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h HeapVal) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapVal) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *HeapVal) Pop() interface{} {
	x := *h
	killallcoreduo := x[len(x)-1]
	x = x[:len(x)-1]
	*h = x
	return killallcoreduo
}
