// Last updated: 11/7/2025, 2:48:00 PM
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left := 1
    right := n
    for left <= right {
        mid := left + (right-left) / 2

        val := guess(mid)
        if val == -1 {
            right = mid - 1
        } else if val == 1 {
            left = mid+1
        } else if val == 0 {
            return mid
        }
    }
    return left
}