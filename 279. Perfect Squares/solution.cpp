class Solution {
public:
    int numSquares(int n) {
        vector<int> squares(n+1, 0);
        for (int num = 1; num <= n; ++num) {
            squares[num] = INT_MAX;
            for (int sq = 1; sq*sq <= num; ++sq) {
                squares[num] = min(squares[num], 1+squares[num-sq*sq]);
            }
        }
        return squares.back();
    }
};
