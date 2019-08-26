class Solution {
public:
    vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
        vector<vector<int>> ans;
        vector<vector<int>> line;
        for (auto b: buildings) {
            line.push_back({b[0], b[2], 0}); // 0 is up
            line.push_back({b[1], b[2], 1}); // 1 is down
        }
        sort(line.begin(), line.end());
        int curr = 0;
        priority_queue<int> height, down;
        for (const auto &b: line) {
            if (b[2] == 0) {
                height.push(b[1]);
                if (b[1] > curr) {
                    curr = b[1];
                    pushAns(ans, {b[0], b[1]});
                }
            } else {
                if (b[1] == curr) {
                    height.pop();
                    while (!down.empty() && !height.empty() && down.top() == height.top()) {
                        down.pop();
                        height.pop();
                    }
                    curr = height.empty() ? 0 : height.top();
                    pushAns(ans, {b[0], curr});
                } else {
                    down.push(b[1]);
                }
            }
        }
        return ans;
    }
    void pushAns(vector<vector<int>> &ans, vector<int> b) {
        if (!ans.empty()) {
            auto last = ans.back();
            if (last[0] == b[0] && last[1] < b[1]) {
                ans.pop_back();
            } else if (last[1] == b[1]) {
                return;
            }
        }
        ans.push_back(b);
    }
};
