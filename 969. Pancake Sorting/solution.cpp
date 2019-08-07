class Solution {
public:
    vector<int> pancakeSort(vector<int>& A) {
        vector<int> ans;
        for (int n = A.size(); n > 0; --n) {
            int idx = 0;
            for (; A[idx] != n; ++idx);
            if (idx == A[idx]-1) continue;
            ans.push_back(idx+1);
            reverse(A.begin(), A.begin()+idx+1);
            ans.push_back(n);
            reverse(A.begin(), A.begin()+n);
        }
        return ans;
    }
};
