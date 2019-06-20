class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t ans = 0;
        for (int idx = 0; idx < 32; ++idx) {
            if (n & uint32_t(1 << idx)) {
                ans |= (1 << (31-idx));
            }
        }
        return ans;
    }
};
