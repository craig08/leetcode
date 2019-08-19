class Solution {
public:
    int rob(vector<int>& nums) {
        int curr = 0, rob2 = 0, rob1 = 0;
        for (int idx = 0; idx < nums.size(); ++idx) {
            curr = max(nums[idx] + rob2, rob1);
            rob2 = rob1;
            rob1 = curr;
        }
        return curr;
    }
};
