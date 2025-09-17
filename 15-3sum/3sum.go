func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    ans := make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
                continue
        } 
        hash := make(map[int]struct{}) 
        for j := i+1; j < len(nums); j++ {
            complete := -nums[j] - nums[i]
            if _, ok := hash[complete]; ok {
                ans = append(ans, []int{nums[i], nums[j], complete})

                for j+1 < len(nums) && nums[j+1]== nums[j] {
                    j++
                }
            
        }
        hash[nums[j]] = struct{}{}
        }
    }
    return ans
}