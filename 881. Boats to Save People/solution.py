class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        i, j, ans = 0, len(people)-1, 0
        while i <= j:
            if people[i]+people[j] <= limit:
                i, j, ans = i+1, j-1, ans+1
            else:
                j, ans = j-1, ans+1
        return ans
