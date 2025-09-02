type SmallestInfiniteSet struct {
	arr *MinHeap
    inHeap map[int]bool
    currMin int
}

func Constructor() SmallestInfiniteSet {
	h := &MinHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{arr: h, inHeap: make(map[int]bool), currMin: 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
    if this.arr.Len() > 0 {
        min := heap.Pop(this.arr).(int)
        delete(this.inHeap, min)
        return min
    }

    res := this.currMin
    this.currMin++
    return res
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.currMin || this.inHeap[num] {
        return
    }

    this.inHeap[num] = true
    heap.Push(this.arr, num)
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
type MinHeap []int

func (m MinHeap) Len() int {
    return len(m)
} 

func (m MinHeap) Less(i,j int) bool {
    return m[i] < m[j] 
}

func (m MinHeap) Swap(i,j int) {
    m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(v any) {
    *m = append(*m, v.(int))
}

func (m *MinHeap) Pop() any {
    old := *m
    n := len(old)
    x := old[n-1]
    *m = old[0 : n-1]
    return x
}