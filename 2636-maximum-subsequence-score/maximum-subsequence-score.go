func maxScore(nums1 []int, nums2 []int, k int) int64 {
	pairs := make([]Pair, len(nums1))
	for i := range pairs {
		pairs[i] = Pair{nums1[i], nums2[i]}
	}

	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i].num2 > pairs[j].num2
	})

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	var currSum int64
	var maxScore int64

	for _, p := range pairs {
		heap.Push(minHeap, p.num1)
		currSum += int64(p.num1)

		if minHeap.Len() > k {
			val := heap.Pop(minHeap).(int)
			currSum -= int64(val)
		}

		if minHeap.Len() == k {
			score := int64(p.num2) * currSum
			if maxScore < score {
				maxScore = score
			}
		}
	}

	return maxScore
}

type Pair struct {
	num1 int
	num2 int
}

type MinHeap []int

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MinHeap) Len() int {
	return len(m)
}

func (m *MinHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	val := old[len(old)-1]
	*m = old[:len(old)-1]
	return val
}
