class Solution {
public:
    int lastStoneWeight(vector<int>& stones) {
        priority_queue<int> heap;
        for (const int stone: stones) {
            heap.push(stone);
        }
        while (heap.size() >= 2) {
            int stone1 = heap.top();
            heap.pop();
            int stone2 = heap.top();
            heap.pop();
            if (stone1 != stone2) {
                heap.push(abs(stone1-stone2));
            }
        }
        return heap.empty() ? 0 : heap.top();
    }
};
