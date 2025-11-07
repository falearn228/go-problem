// Last updated: 11/7/2025, 2:46:37 PM
func maxVowels(s string, k int) int {
    currVowel := 0

    vowels := []byte{ 'a', 'e', 'i', 'o', 'u'}

    for i := range k {
        if slices.Contains(vowels, s[i]) {
            currVowel++
        }
    }

    maxVowel := currVowel
    for i := k; i < len(s); i++ {
        if slices.Contains(vowels, s[i]) {
            currVowel++
        }
        if slices.Contains(vowels, s[i-k]) {
            currVowel--
        }
        maxVowel = max(currVowel, maxVowel)
    } 
    return maxVowel
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}