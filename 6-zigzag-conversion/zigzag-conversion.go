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

    ans := ""
    for i := range rows {
        ans += rows[i].String()
    }

    return ans
}