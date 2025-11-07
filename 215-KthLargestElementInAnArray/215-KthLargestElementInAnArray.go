// Last updated: 11/7/2025, 2:48:17 PM
func findKthLargest(nums []int, k int) int {
    heap := &heap{}
    for i := range nums {
        if len(heap.arr) < k {
            heap.Push(nums[i])
        } else if heap.arr[0] < nums[i] {
            heap.Pop()
            heap.Push(nums[i])
        }
    }
    return heap.arr[0]
}

type heap struct{
    arr []int
}

func (h* heap) Push(value int) {
    if len(h.arr) == 0 {
        h.arr = append(h.arr, value)
        return
    }
    
    h.arr = append(h.arr, value)
    h.Up(len(h.arr)-1)
}
// 2*i +1 -> 7, 8 (3) -> 
func (h* heap) Up(ind int) {
    for ind > 0 {
        p := (ind-1) / 2

        if h.arr[ind] >= h.arr[p] {
            return
        }
        h.arr[ind], h.arr[p] = h.arr[p], h.arr[ind]
        ind = p
    }
} 

func (h* heap) Down(i int) {
    n := len(h.arr)
    for {
        left := 2*i+1
        right := 2*i+2
        best := i

        if left < n && h.arr[best] > h.arr[left] {
            best = left
        } 
        if right < n && h.arr[best] > h.arr[right] {
            best = right
        } 
        if best == i {
            break
        }
        h.arr[i], h.arr[best] = h.arr[best], h.arr[i]
        i = best
    }
} 

func (h *heap) Pop() (int, bool) {
	if len(h.arr) == 0 {
        return 0, false
    }

    n := len(h.arr) - 1
    h.arr[0], h.arr[n] = h.arr[n], h.arr[0]
    val := h.arr[n]
    h.arr = h.arr[:n]

    if len(h.arr) > 0 {
        h.Down(0)
    }

    return val, true
}
