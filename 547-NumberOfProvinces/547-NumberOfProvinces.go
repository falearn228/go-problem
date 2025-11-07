// Last updated: 11/7/2025, 2:47:46 PM
func findCircleNum(isConnected [][]int) int {
    visited := make([]bool , len(isConnected))
    
    provincies := 0

    for i := 0; i < len(isConnected); i++ {
        if !visited[i] {
            provincies++
            stack := []int{i}
            for len(stack) > 0 {
            room := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if visited[room] {
                continue
            }
            visited[room] = true

            tops := isConnected[room]
                for province, conn := range tops {
                    if conn == 1 && !visited[province] {
                        stack = append(stack, province)
                    }
                }
            }
        }
    }

    return provincies
}