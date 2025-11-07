// Last updated: 11/7/2025, 2:47:58 PM
type RandomizedSet struct {
    setInd map[int]int // value -> index
    arr []int // value repo
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        setInd: make(map[int]int),
        arr: make([]int, 0, 1024),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.setInd[val]; ok {
        return false
    } 
    this.arr = append(this.arr, val)
    index := len(this.arr) - 1
    this.setInd[val] = index
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    index, ok := this.setInd[val]
    if !ok {
        return false
    }
    

    lastValue := this.arr[len(this.arr)-1]
    this.arr[index] = lastValue
    this.setInd[lastValue] = index 
    delete(this.setInd, val)
    this.arr = this.arr[:len(this.arr)-1]

    return true
}


func (this *RandomizedSet) GetRandom() int {
    randomIndex := rand.Intn(len(this.arr))
    return this.arr[randomIndex]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */