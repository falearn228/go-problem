func canCompleteCircuit(gas []int, cost []int) int {
    var totalTank, currTank, i, currStation int

    n := len(gas)
    for i < n {
        diff := gas[i] - cost[i]
        totalTank += diff
        currTank += diff
        if currTank < 0 { // ушли в минус по бензину, недоживем до следующей станции 
            // обновляем станцию, станции ниже i не могут быть стартовыми
            currStation = i + 1
            currTank = 0
        }
        i++
    } 

    if totalTank < 0 {
        return -1
    }
    return currStation
}