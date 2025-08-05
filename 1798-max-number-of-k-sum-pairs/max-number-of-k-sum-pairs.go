func maxOperations(nums []int, k int) int {
    kSet := make(map[int]int)
    opers := 0

    for _, v := range nums {
        diff := k - v
        val := kSet[diff]
        if val <= 0  {
            kSet[v]++
            continue
        }

        kSet[diff]--
        opers++

        
    }
    return opers
}