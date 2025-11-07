// Last updated: 11/7/2025, 2:49:06 PM

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	n := len(words)
	i := 0
	for i < n {
		lineWords := []string{}
		currLength := 0
		for i < n && len(lineWords)+len(words[i])+currLength <= maxWidth {
			lineWords = append(lineWords, words[i])
			currLength += len(words[i])
			i++
		}

		var sb strings.Builder
		isLastLine := i == n
		isOneWordLine := len(lineWords) == 1

		if isLastLine || isOneWordLine {
			sb.WriteString(strings.Join(lineWords, " "))
			sb.WriteString(strings.Repeat(" ", maxWidth-sb.Len()))
		} else {
			totalSpaces := maxWidth - currLength
			gaps := len(lineWords) - 1
			spacesPerGap := totalSpaces / gaps
			extraSpaces := totalSpaces % gaps

			for j := range lineWords {
				sb.WriteString(lineWords[j])
				if j < gaps {
					sb.WriteString(strings.Repeat(" ", spacesPerGap))
					if j < extraSpaces {
						sb.WriteString(" ")
					}
				}

			}
		}
		res = append(res, sb.String())

	}
	return res
}
