func canConstruct(ransomNote string, magazine string) bool {
    redis := make([]int, 26)
    for i := range magazine {
        redis[magazine[i]-'a']++
    }

    for i := range ransomNote {
        cnt := redis[ransomNote[i]-'a']
        if cnt <= 0 {
            return false
        }
        redis[ransomNote[i]-'a']--
    }

    return true
}