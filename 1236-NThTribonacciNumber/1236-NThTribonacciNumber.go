// Last updated: 11/7/2025, 2:47:00 PM
func tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 || n == 2 {
        return 1
    }
    a := 0
    b := 1
    c := 1
    var res int
    // tribb[3] = ...tribb[2,1,0]

    for i := 3; i <= n; i++ {
        res = a + b + c
        a, b, c = b, c, res
    }

    return res
}