// Last updated: 11/7/2025, 2:47:16 PM
func canVisitAllRooms(rooms [][]int) bool { 
    visited := make(map[int]bool)
    stack := make([]int, 1)
    visitedCount := 0
    for len(stack) > 0 {
        room := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if visited[room] {
            continue
        }
        visited[room] = true
        visitedCount++

        tops := rooms[room]
        for _, child := range tops {
            if !visited[child] {
                stack = append(stack, child)
            }
        }
    }
    
    return visitedCount == len(rooms)
}