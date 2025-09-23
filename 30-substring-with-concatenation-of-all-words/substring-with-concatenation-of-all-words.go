func findSubstring(s string, words []string) []int {
    ans := make([]int, 0, len(words))
    word_count := make(map[string]int, len(words))
    for _, word := range words {
        word_count[word]++ 
    }
    wordLen := len(words[0])
	numWords := len(words)
	substringLen := wordLen * numWords

    if substringLen > len(s) {
        return ans
    }

    for i := 0; i < wordLen; i++ {
        left := i
        count := 0
        window := make(map[string]int)

        for right := i; right <= len(s)-wordLen; right += wordLen {
            word := s[right : right+wordLen]

            if targetCount, ok := word_count[word]; ok {
                window[word]++
                count++

                for window[word]> targetCount {
                    leftWord := s[left : left+wordLen]
                    window[leftWord]--
                    count--
                    left += wordLen
                }

                if count == numWords {
                    ans = append(ans, left)
                    leftWord := s[left : left+wordLen]
                    window[leftWord]--
                    count--
                    left += wordLen
                }
            } else {
                clear(window)
                count = 0
                left = right + wordLen
            }
        }
    }

    return ans
}