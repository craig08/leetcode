class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        if (grid.empty()) return 0;
        vector<vector<int>> group(grid.size(), vector<int>(grid[0].size(), 0));
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[0].size(); ++j) {
                if (grid[i][j] == '0') continue;
                group[i][j] = i*group[0].size() + j;
                if (i > 0 && grid[i-1][j] != '0') {
                    unionIsland(grid, group, i, j, i-1, j);
                }
                if (j > 0 && grid[i][j-1] != '0') {
                    unionIsland(grid, group, i, j, i, j-1);
                }
            }
        }
        unordered_set<int> islands;
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[0].size(); ++j) {
                if (grid[i][j] != '0') {
                    islands.insert(find(grid, group, i, j));
                }
            }
        }
        return islands.size();
    }
    void unionIsland(const vector<vector<char>> &grid, vector<vector<int>> &group, int i1, int j1, int i2, int j2) {
        int p1 = find(grid, group, i1, j1), p2 = find(grid, group, i2, j2);
        if (p1 == p2) return;
        if (p1 < p2) {
            group[p2/group[0].size()][p2%group[0].size()] = p1;
        } else {
            group[p1/group[0].size()][p1%group[0].size()] = p2;
        }
    }
    int find(const vector<vector<char>> &grid, vector<vector<int>> &group, int i, int j) {
        if (i*group[0].size()+j == group[i][j]) return group[i][j];
        return group[i][j] = find(grid, group, group[i][j]/group[0].size(), group[i][j]%group[0].size());
    }
};
