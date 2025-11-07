// Last updated: 11/7/2025, 2:46:47 PM
func suggestedProducts(products []string, searchWord string) [][]string {
    res := make([][]string, 0)

    sort.Strings(products)

    var pref string
    currSugg := products
    for _, char := range searchWord {
        pref += string(char)

        nextSugg := []string{}
        for _, product := range currSugg {
            if strings.HasPrefix(product, pref) {
                nextSugg = append(nextSugg, product)
            }
        }

        if len(nextSugg) > 3{
            res = append(res, nextSugg[:3])
        } else {
            res = append(res, nextSugg)
        }
        currSugg = nextSugg
    }

    return res
}