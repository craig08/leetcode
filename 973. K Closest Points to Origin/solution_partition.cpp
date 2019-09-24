class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int K) {
        int left = 0, right = points.size();
        while (left < right) {
            int idx = partition(points, left, right);
            if (idx == K-1) {
                break;
            } else if (idx < K-1) {
                left = idx+1;
            } else {
                right = idx;
            }
        }
        points.resize(K);
        return points;
    }
    int partition(vector<vector<int>> &points, int start, int end) {
        int pivot = start, pivotDis = dis(points[end-1]);
        for (int idx = start; idx < end-1; ++idx) {
            if (dis(points[idx]) < pivotDis) {
                swap(points[idx], points[pivot++]);
            }
        }
        swap(points[pivot], points[end-1]);
        return pivot;
    }
    int dis(const vector<int> &point) {
        return point[0]*point[0] + point[1]*point[1];
    }
};
