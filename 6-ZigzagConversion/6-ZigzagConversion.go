// Last updated: 11/7/2025, 2:49:41 PM
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    rows := make([]strings.Builder, numRows)
    
    currBld := 0
    direction := 1
    for i := range s {
        rows[currBld].WriteByte(s[i])

        if currBld >= numRows-1 {
            direction = -1
        } else if currBld <= 0 {
            direction = 1
        } 

        if direction == 1 {
            currBld++
        } else {
            currBld--
        }
    }

    var finalResult strings.Builder
    for i := range rows {
        finalResult.WriteString(rows[i].String())
    }

    return finalResult.String()
}