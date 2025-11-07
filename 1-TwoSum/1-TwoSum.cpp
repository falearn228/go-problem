// Last updated: 11/7/2025, 2:49:51 PM
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> mp; vector<int> res;
        for(int i = 0; i < nums.size(); i++) {
            if(mp.find(target - nums[i]) != mp.end()) {
                res.push_back(mp[target-nums[i]]);
                res.push_back(i);
                return res;
            }
            mp[nums[i]] = i;
        }
        return res;
    }
};