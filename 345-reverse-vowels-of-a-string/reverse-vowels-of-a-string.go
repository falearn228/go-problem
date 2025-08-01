func reverseVowels(s string) string {
    if len(s) == 1 {
        return s
    }

	i, j := 0, len(s)-1

	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	result := make([]byte, len(s))
	for i <= j {
		if !slices.Contains(vowels, s[i]) {
			result[i] = s[i]
			i++
		}
		if !slices.Contains(vowels, s[j]) {
			result[j] = s[j]
			j--
		}
		if slices.Contains(vowels, s[i]) && slices.Contains(vowels, s[j]) {
			result[i] = s[j]
			result[j] = s[i]
			fmt.Println(string(s[j]), string(s[i]))
			i++
			j--

		}

	}

	return string(result)
}