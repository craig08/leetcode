class Solution {
public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        wordList.push_back(beginWord);
        vector<bool> used(wordList.size(), false);
        queue<int> path;
        path.push(wordList.size()-1);
        used[wordList.size()-1] = true;
        int step = 2;
        while (!path.empty()) {
            for (int idx = path.size(); idx > 0; --idx) {
                int node = path.front();
                path.pop();
                for (int adj = 0; adj < wordList.size(); ++adj) {
                    if (!used[adj] && adjacent(wordList[node], wordList[adj])) {
                        if (wordList[adj] == endWord) {
                            return step;
                        }
                        used[adj] = true;
                        path.push(adj);
                    }
                }
            }
            ++step;
        }
        return 0;
    }
    bool adjacent(const string &s1, const string &s2) {
        if (s1.size() != s2.size()) return false;
        int diff = 0;
        for (int idx = 0; idx < s1.size(); ++idx) {
            if (s1[idx] != s2[idx] && ++diff > 1) {
                return false;
            }
        }
        return diff == 1;
    }
};
