func isHappy(n int) bool {
    prev := make(map[int]bool)
    for prev[n] != true {
        if n == 1 {
            return true
        }
        // n % 10 = 9   ->  // 10 ...
        sum := 0
        prev[n] = true
        for n != 0 {
            sum += (n%10) * (n%10)
            n /= 10 
        }
        n = sum
    }

    return false
}