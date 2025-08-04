func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

    writePtr := 0
    readPtr := 0

    for readPtr < len(chars) {
        currentChar := chars[readPtr]
        count := 0

        for readPtr < len(chars)  && chars[readPtr] == currentChar {
            count++
            readPtr++
        }

        chars[writePtr] = currentChar
        writePtr++

        if count > 1 {
            countStr := strconv.Itoa(count)
            for _, symbol := range countStr {
                chars[writePtr] = byte(symbol)
                writePtr++
            }
        }
    }
 
	return writePtr
}