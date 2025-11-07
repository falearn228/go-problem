// Last updated: 11/7/2025, 2:47:11 PM
type RecentCounter struct {
    reqs []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        reqs: make([]int, 0),
    }
}

const T = 3000

func (this *RecentCounter) Ping(t int) int {
    if len(this.reqs) == 0 {
        this.reqs = append(this.reqs, t)
        return 1
    }

    this.reqs = append(this.reqs, t)
    for this.reqs[0] < t-T && this.reqs[0] < t {
       this.reqs = this.reqs[1:]
    }
    return len(this.reqs)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */