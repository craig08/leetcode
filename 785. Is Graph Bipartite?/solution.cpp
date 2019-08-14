class Solution {
public:
    bool isBipartite(vector<vector<int>>& graph) {
        vector<int> group(graph.size(), 0);
        for (int node = 0; node < graph.size(); ++node) {
            if (group[node] == 0 && !traverse(graph, group, node, 1)) {
                return false;
            }
        }
        return true;
    }
    bool traverse(const vector<vector<int>> &graph, vector<int> &group, int node, int groupNum) {
        if (group[node] != 0) {
            return group[node] == groupNum;
        }
        group[node] = groupNum;
        for (int n = 0; n < graph[node].size(); ++n) {
            if (!traverse(graph, group, graph[node][n], 1 + groupNum%2)) {
                return false;
            }
        }
        return true;
    }
};
