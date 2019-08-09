class Solution {
public:
    void sortColors(vector<int>& nums) {
        int i = 0, j = 0, n = nums.size()-1;
        while (j <= n) {
            if (nums[j] < 1) {
                swap(nums[i++], nums[j++]);
            } else if (nums[j] > 1) {
                swap(nums[j], nums[n--]);
            } else {
                ++j;
            }
        }
    }
};
