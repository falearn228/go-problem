// Last updated: 11/7/2025, 2:46:11 PM
func minimumSum(num int) int {
    arr := make([]int, 0)
    for num > 0 {
        arr = append(arr, num % 10)
        num /= 10 
    } 

    slices.Sort(arr)

    num1 := 0
    num2 := 0

    i, j := 0, 1
    for i < len(arr) && j < len(arr) {
        num1 *= 10
        num1 += arr[i]

        num2 *= 10
        num2 += arr[j]
        i += 2
        j += 2
    }

    return num1 + num2
}