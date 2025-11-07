// Last updated: 11/7/2025, 2:49:42 PM
class Solution {
public:
    std::string longestPalindrome(std::string s) {
std::string extendedString = "#"; // Разделитель символов для обработки четных палиндромов
    for (char c : s) {
        extendedString += c;
        extendedString += '#';
    }

    int n = extendedString.length();
    std::vector<int> palindromeLengths(n, 0);
    int center = 0, right = 0;

    for (int i = 0; i < n; i++) {
        int mirror = 2 * center - i;

        if (i < right) {
            palindromeLengths[i] = std::min(right - i, palindromeLengths[mirror]);
        }

        int a = i + (1 + palindromeLengths[i]);
        int b = i - (1 + palindromeLengths[i]);

        while (a < n && b >= 0 && extendedString[a] == extendedString[b]) {
            palindromeLengths[i]++;
            a++;
            b--;
        }

        if (i + palindromeLengths[i] > right) {
            center = i;
            right = i + palindromeLengths[i];
        }
    }

    int maxLen = 0, maxCenter = 0;
    for (int i = 0; i < n; i++) {
        if (palindromeLengths[i] > maxLen) {
            maxLen = palindromeLengths[i];
            maxCenter = i;
        }
    }

    // Извлекаем палиндром из расширенной строки
    int start = (maxCenter - maxLen) / 2;
    return s.substr(start, maxLen);
}
};