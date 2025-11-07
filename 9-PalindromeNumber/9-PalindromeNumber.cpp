// Last updated: 11/7/2025, 2:49:39 PM
class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0 || x != 0 && x % 10 == 0) return false;
        long res = 0;
        while (x > res) {
            res = res*10 + x %10;
            x /= 10;    
        }
        return (x == res || x == res / 10) ? 1 : 0;
    }
};