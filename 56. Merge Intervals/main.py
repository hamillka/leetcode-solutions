class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: x[0])
        merged = []

        prev = intervals[0]

        for itr in intervals[1:]:
            if itr[0] <= prev[1]:
                prev[1] = max(prev[1], itr[1])
            else:
                merged.append(prev)
                prev = itr
        
        merged.append(prev)

        return merged

