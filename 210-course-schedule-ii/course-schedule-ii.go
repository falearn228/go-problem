func findOrder(numCourses int, prerequisites [][]int) []int {
    if numCourses == 1 {
        return []int{0}
    }

    graph := make([][]int, numCourses)
    for i := range prerequisites {
        want, need := prerequisites[i][0], prerequisites[i][1]
        graph[need] = append(graph[need], want) 
    }

    res := make([]int, 0, numCourses)
    state := make([]int, numCourses)

    var dfs func(int) bool
    dfs = func(curr int) bool {
        if state[curr] == 2 {
            return true
        } 

        if state[curr] == 1 {
            return false
        }

        state[curr] = 1
        for _, nei := range graph[curr] {
            // need -> want, want, want...
            if !dfs(nei) {
                return false
            }
        }

        state[curr] = 2
        res = append(res, curr)
        return true 
    }

    for crs := range numCourses {
        if state[crs] == 0 {
            if !dfs(crs) {
                return []int{}
            }
        }
    }

    slices.Reverse(res)
    return res
}