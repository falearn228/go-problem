type MinStack struct {
	stack []Pair
}

type Pair struct {
	value int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]Pair, 0, 100),
	}
}

func (this *MinStack) Push(val int) {
	minval := val
	if len(this.stack) > 0 {
		// min with curr val and prev val
		minval = min(this.stack[len(this.stack)-1].min, val)
	}
	this.stack = append(this.stack, Pair{val, minval})
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].value
}

func (this *MinStack) GetMin() int {
	res := this.stack[len(this.stack)-1].min
    // this.stack = this.stack[:len(this.stack)-1]
    return res
}