// Last updated: 11/7/2025, 2:49:40 PM
class Solution {
public:
    int reverse(int x) {
          if (x == 0) {
            return 0;
        }

        bool isNegative = (x < 0);
        long long absX = std::abs(static_cast<long long>(x)); // Handle potential overflow

        // Convert the number to a string
        std::string strX = std::to_string(absX);

        // Reverse the string
        std::reverse(strX.begin(), strX.end());

        // Convert the reversed string back to an integer
        long long result = stoll(strX);

        // Handle overflow by checking the result against INT_MAX and INT_MIN
        if (result > INT_MAX || result < INT_MIN) {
            return 0;
        }

        // Apply the original sign
        return isNegative ? -result : result;
    }
};