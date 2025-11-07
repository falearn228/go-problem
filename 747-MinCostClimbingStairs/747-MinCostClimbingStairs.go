// Last updated: 11/7/2025, 2:47:21 PM
func minCostClimbingStairs(cost []int) int {
    var stepZero int
    var stepOnce int

    var finalStep int
    for i := 1; i < len(cost); i++ {
        finalStep = min(stepOnce+ cost[i], stepZero+ cost[i-1])
        stepZero = stepOnce
        stepOnce = finalStep
    }

    return finalStep
 }