// Last updated: 11/7/2025, 2:49:48 PM
#include <vector>
#include <iostream>
#include <algorithm> // Добавленный заголовок для find
#include <string>    // Добавленный заголовок для std::string

class Solution {
public:
    int lengthOfLongestSubstring(std::string s) {
        std::vector<char> dumb;
        int maxLength = 0;

        for (size_t i = 0; i < s.size(); i++) {
            auto it = std::find(dumb.begin(), dumb.end(), s[i]);

            if (it != dumb.end()) {
                // Найден повторяющийся символ, обновляем максимальную длину
                maxLength = std::max(maxLength, static_cast<int>(dumb.size()));
                // Удаляем все элементы до найденного символа, включительно
                dumb.erase(dumb.begin(), it + 1);
            }

            // Добавляем текущий символ в вектор
            dumb.push_back(s[i]);
        }

        // Проверка максимальной длины после выхода из цикла
        maxLength = std::max(maxLength, static_cast<int>(dumb.size()));
        
        return maxLength;
    }
};
