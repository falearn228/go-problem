// Last updated: 11/7/2025, 2:49:32 PM
class Solution {
public:
    bool isValid(string s) {
        if (s.size() == 1) return false; 
        std::vector <char> res;
        std::string f = ")}]";
        for (char& sym : s) {
            res.push_back(sym); 
            switch (sym) {
            case ']':
                if (res.size() >= 2 && res[res.size() - 2] == '[') {
                    res.pop_back();
                    res.pop_back();
                }
                break;
            case ')':
                 if (res.size() >= 2 && res[res.size() - 2] == '(') {
                    res.pop_back();
                    res.pop_back();
                }
                break;
            case '}':
                 if (res.size() >= 2 && res[res.size() - 2] == '{') {
                    res.pop_back();
                    res.pop_back();
                }
                break;
            }
        }
        return res.empty();
    }
};

/*

res [(((())))]
*/