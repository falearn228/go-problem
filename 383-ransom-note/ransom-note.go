func canConstruct(ransomNote string, magazine string) bool {
    redis := make(map[byte]int, len(magazine))
    for i := range magazine {
        redis[magazine[i]]++
    }

    for i := range ransomNote {
        cnt, ok := redis[ransomNote[i]]
        if !ok {
            return false
        }
        redis[ransomNote[i]]--
        if cnt-1 < 0 {
            return false
        }
    }

    return true
}