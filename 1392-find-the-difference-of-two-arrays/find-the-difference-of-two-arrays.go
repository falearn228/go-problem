func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]struct{})
    set2 := make(map[int]struct{})

    for i := 0; i < len(nums1); i++ {
        set1[nums1[i]] = struct{}{}
    }
    for i := 0; i < len(nums2); i++ {
        set2[nums2[i]] = struct{}{}
    }

    result := make([][]int, 2)
    for key := range set1 {
        if _, ok := set2[key]; !ok {
            result[0] = append(result[0], key)
            delete(set1, key)
        }
    }
     for key := range set2 {
        if _, ok := set1[key]; !ok {
            result[1] = append(result[1], key)
            delete(set2, key)
        }
    }
    return result
}