
// Для заданного целого числа n верните массив ans длиной n + 1, в котором для каждого i (0 <= i <= n) ans[i] — это количество единиц в двоичном представлении i.
func countBits(n int) []int {
    ans := make([]int, n+1)
    
    for i := range ans {
        ans[i] = ans[i >> 1] + (i & 1)
    }
    return ans
}