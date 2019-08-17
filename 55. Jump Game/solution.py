class Solution:
    def canJump(self, nums: List[int]) -> bool:
        jump, idx = 0, 0
        while jump < len(nums)-1 and idx <= jump:
            jump = max(jump, idx + nums[idx])
            idx += 1
        return jump >= len(nums)-1
