func tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 || n == 2 {
        return 1
    }
    tribb := make([]int, n+1)
    tribb[0] = 0
    tribb[1] = 1
    tribb[2] = 1
    // tribb[3] = ...tribb[2,1,0]

    for i := 3; i <= n; i++ {
        tribb[i] = tribb[i-1] + tribb[i-2] + tribb[i-3]
    }

    return tribb[n]
}