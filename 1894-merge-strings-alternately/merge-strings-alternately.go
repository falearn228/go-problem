func mergeAlternately(word1 string, word2 string) string {
	i1 := 0
	i2 := 0
	var mergedStr strings.Builder
	mergedStr.Grow(len(word1) + len(word2))
	for i1 < len(word1) && i2 < len(word2) {
		if i1 == i2 {
			mergedStr.WriteByte(word1[i1])
			i1++
		} else {
			mergedStr.WriteByte(word2[i2])
			i2++
		}
	}

	for i1 < len(word1) {
		mergedStr.WriteByte(word1[i1])
		i1++
	}
	for i2 < len(word2) {
		mergedStr.WriteByte(word2[i2])
		i2++
	}

	return mergedStr.String()
}
