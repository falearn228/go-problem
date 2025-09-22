func findSubstring(s string, words []string) []int {
	// Проверяем крайние случаи: если строка, срез слов или сами слова пустые.
	if len(s) == 0 || len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	numWords := len(words)
	substringLen := wordLen * numWords
	
	// Если искомая подстрока длиннее основной строки, решения нет.
	if substringLen > len(s) {
		return []int{}
	}

	// 1. Создаем карту для подсчета частоты каждого слова в `words`.
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	result := make([]int, 0)

	// 2. Проходим по строке `s` с помощью "скользящего окна".
	// Окно будет иметь размер `substringLen`. Мы сдвигаем его на 1 символ.
	for i := 0; i <= len(s)-substringLen; i++ {
		// 3. Для каждого окна создаем свою карту `seen` для подсчета слов.
		seen := make(map[string]int)
		
		// `j` - счетчик слов внутри текущего окна.
		j := 0
		for ; j < numWords; j++ {
			// Вырезаем слово из строки `s` в текущей позиции окна.
			wordStart := i + j*wordLen
			word := s[wordStart : wordStart+wordLen]

			// Проверяем, есть ли такое слово в нашей исходной карте `wordCount`.
			if count, ok := wordCount[word]; ok {
				// Если слово есть, увеличиваем его счетчик в карте `seen`.
				seen[word]++
				// Если мы встретили это слово больше раз, чем нужно,
				// то это окно нам не подходит, прерываем внутренний цикл.
				if seen[word] > count {
					break
				}
			} else {
				// Если слова нет в `wordCount`, окно неверное, прерываемся.
				break
			}
		}

		// 4. Если внутренний цикл успешно прошел все слова (`j == numWords`),
		// значит, мы нашли правильную комбинацию. Добавляем индекс `i` в результат.
		if j == numWords {
			result = append(result, i)
		}
	}

	return result
}