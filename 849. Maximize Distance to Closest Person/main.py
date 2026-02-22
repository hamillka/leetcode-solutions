class Solution:
    def maxDistToClosest(self, seats: List[int]) -> int:
        max_dist = 0
        prev = -1
        
        for i in range(len(seats)):
            if seats[i] == 1:
                if prev == -1:
                    max_dist = i
                else:
                    max_dist = max(max_dist, (i - prev) // 2)
                prev = i
        
        max_dist = max(max_dist, len(seats) - 1 - prev)
        
        return max_dist
