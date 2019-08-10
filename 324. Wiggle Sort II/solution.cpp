class Solution {
public:
    int N;
    void wiggleSort(vector<int>& nums) {
        // find median
        N = nums.size();
        nth_element(nums.begin(), nums.begin() + N/2, nums.end());
        int med = nums[N/2];
        int i = 0, j = 0, n = N-1;
        while (j <= n) {
            if (nums[idxMap(j)] < med) {
                swap(nums[idxMap(i++)], nums[idxMap(j++)]);
            } else if (nums[idxMap(j)] > med) {
                swap(nums[idxMap(j)], nums[idxMap(n--)]);
            } else {
                ++j;
            }
        }
    }
    /* To avoid errors, an index mapping must keep the numbers equal
       to median far from each other. We can imply that the same order
       should be applied to both halves. In addition, considering the
       case with only 4 elements: 0 1 2 3, 1 and 2 are medians, so
       0 2 1 3 will get error. If we use descending order to arrange,
       we can get 1 3 0 2, making the tail of first half and the head
       of second half as the head and tail of the whole array. */
    int idxMap(int i) {
        if (i >= (N+1)/2) {
            return N*2 - 1 - i*2;
        }
        if (N & 1) {
            return N - 1 - i*2;
        }
        return N - 2 - i*2;
    }
};
