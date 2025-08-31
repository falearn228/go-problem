func suggestedProducts(products []string, searchWord string) [][]string {
    res := make([][]string, len(searchWord))

    sort.Strings(products)

    pref := make([]byte, 0)
    for i := 0; i < len(searchWord); i++ {
        pref = append(pref, searchWord[i])
        for j := range len(products) {
            if len(res[i]) == 3 {
                break
            }
            if strings.HasPrefix(products[j], string(pref)) {
                res[i] = append(res[i], products[j])
            }
        }
    }

    return res
}