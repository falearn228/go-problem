func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    indMap := make(map[int]struct{})

    for i := range nums {
        indMap[nums[i]] = struct{}{}
    }

    answer := 0
    for value := range indMap {
        if _, found := indMap[value-1]; found {
            continue
        }

        
        currNum := value
        currLen := 1

        for {
            if _, found := indMap[currNum+1]; found {
                currLen++
                currNum++
            } else {
                break
            }
        }
        if currLen > answer {
            answer = currLen
        }

    }
    return answer
}