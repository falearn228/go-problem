func isIsomorphic(s string, t string) bool {
    kafka := make(map[byte]byte, len(s))
    rabbitmq := make(map[byte]byte, len(t))

    for i := range s {
        char_s, okS := kafka[s[i]]
        char_t, okT := rabbitmq[t[i]]
        if okS || okT {
            if t[i] != char_s || s[i] != char_t {
                return false
            }
        }
        kafka[s[i]] = t[i]
        rabbitmq[t[i]] = s[i]
    } 

    return true
}