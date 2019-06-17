class Solution {
public:
    vector<int> getRow(int rowIndex) {
        vector<int> arr(rowIndex+1, 1);
        for (int row = 1; row < rowIndex; ++row) {
            int last = 1;
            for (int idx = 1; idx <= row; ++idx) {
                int curr = arr[idx];
                arr[idx] += last;
                last = curr;
            }
        }
        return arr;
    }
};
