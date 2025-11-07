// Last updated: 11/7/2025, 2:47:17 PM

func numTilings(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    shapes := make([]int, n+1)
    shapes[0] = 1
    shapes[1] = 1
    shapes[2] = 2
    
    for i := 3; i <= n; i++ {
        shapes[i] = (2*shapes[i-1] % MOD + shapes[i-3]%MOD) % MOD 
    }
    

    return shapes[n]
}

const MOD = 1000000007