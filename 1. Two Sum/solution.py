class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        idxs = dict()
        for idx, num in enumerate(nums):
            if (target - num) in idxs:
                return [idxs[target-num], idx]
            idxs[num] = idx
        return []
