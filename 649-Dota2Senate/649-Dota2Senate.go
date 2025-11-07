// Last updated: 11/7/2025, 2:47:28 PM
func predictPartyVictory(senate string) string {
    queueR := make([]int, 0, len(senate))
    queueD := make([]int, 0, len(senate))
    n := len(senate)
    for i := range senate {
        if senate[i] == 'R' {
            queueR = append(queueR, i)
        } else {
            queueD = append(queueD, i)
        }
    }
    for len(queueR) > 0 && len(queueD) > 0 {
        if queueR[0] < queueD[0] {
            queueR = append(queueR, queueR[0]+n)
            queueR = queueR[1:]
            queueD = queueD[1:]
            
        } else {
            queueD = append(queueD, queueD[0]+n)
            queueR = queueR[1:]
            queueD = queueD[1:]
        }
    }

    if len(queueR) == 0 {
        return "Dire"
    } else {
        return "Radiant"
    }
}

// 