// Last updated: 11/7/2025, 2:49:44 PM
class Solution {
public:
    string convert(string s, int numRows) {
          if (numRows == 1 || numRows >= s.size()) {
            return s;  // No conversion needed
        }

        std::vector<std::vector<char>> map(numRows, std::vector<char>());
        int count = 0;
        int direction = 1; // 1 for down, -1 for up

        for (size_t j = 0; j < s.size(); j++) {
            map[count].push_back(s[j]);

            if (count == 0) {
                direction = 1;  // Change direction to down when reaching the top row
            } else if (count == numRows - 1) {
                direction = -1; // Change direction to up when reaching the bottom row
            }

            count += direction;
        }

        std::string result;
        for (const auto &row : map) {
            result += std::string(row.begin(), row.end());
        }

        return result;
    }
};