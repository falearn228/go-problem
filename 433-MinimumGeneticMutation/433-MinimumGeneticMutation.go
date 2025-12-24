// Last updated: 12/24/2025, 12:17:50 PM
type Node struct {
    gen string
    steps int
}

func minMutation(startGene string, endGene string, bank []string) int {
    if startGene == endGene {
        return 0
    }

    visited := make(map[string]bool, len(bank)+1)
    visited[startGene] = true
    q := []Node{{gen: startGene, steps: 0}}

    for len(q) > 0 {
        curr := q[0]
        q = q[1:]

        if curr.gen == endGene {
            return curr.steps
        }

        for _, v := range bank {

            if visited[v] {
                continue
            }
            
            diff := 0
            for ind := range v {
                if v[ind] != curr.gen[ind] {
                    diff++
                  
                    if diff > 1 {
                        break 
                    }
                }
            }

            if diff == 1 {
                visited[v] = true 
                q = append(q, Node{gen: v, steps: curr.steps + 1})
            }
        }
    }

    return -1
}
