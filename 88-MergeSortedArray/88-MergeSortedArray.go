// Last updated: 11/7/2025, 2:49:00 PM
func merge(nums1 []int, m int, nums2 []int, n int)  {
    if n == 0 {
        return
    }
    i1, i2, i3 := m - 1, n - 1, m + n - 1
    
    for i2 >= 0 {
        if i1 >= 0 && nums1[i1] > nums2[i2] {
            nums1[i3] = nums1[i1]
            i1--
        } else {
            nums1[i3] = nums2[i2]
            i2--
        }
        i3--
    } 

}