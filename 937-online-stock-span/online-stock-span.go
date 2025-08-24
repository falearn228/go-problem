type pair struct {
    span int
    price int
}

type StockSpanner struct {
    stack []pair
}


func Constructor() StockSpanner {
    return StockSpanner{
        stack: make([]pair, 0, 1024),
    }
}


func (this *StockSpanner) Next(price int) int {
    span := 1
    for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1].price {
        span += this.stack[len(this.stack)-1].span
        this.stack = this.stack[:len(this.stack)-1]
    }
    
    this.stack = append(this.stack, pair{span: span, price: price})
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */