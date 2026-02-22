class Solution:
    def isReflected(self, points: list[list[int]]) -> bool:
        if not points:
            return True

        min_x = min(p[0] for p in points)
        max_x = max(p[0] for p in points)
        
        s = min_x + max_x
        
        point_set = { (x, y) for x, y in points }

        for x, y in points:
            reflected = (s - x, y)
            if reflected not in point_set:
                return False
       
        return True

'''
[[1, 1] [-1, 1]]

min_x = -1
max_x = 1

s = -1+1 = 0

1. 
refl = (0 - 1, 1) = (-1, 1)

2.
refl = (0 - (-1), 1) = (1, 1)


---
[[1, 1] [2, 2]]

min_x = 1
max_x = 2

s = 1 + 2 = 3

1. 
refl = (3 - 1, 1) = (2, 1)

2.
refl = (3 - 2, 2) = (1, 2)


---
[[1, 1], [2, 1], [-1, 1], [-2, 1]]

min_x = -2
max_x = 2 

s = -2 + 2 = 0 

1.
refl = (0 - 1, 1) = (-1, 1)

2.
refl = (0 - 2, 1) = (-2, 1)

3.
refl = (0 - (-1), 1) = (1, 1)

4.
refl = (0 - (-2), 1) = (2, 1)
'''

if __name__ == '__main__':
    sol = Solution()
    
    # Example 1:
    points1 = [[1, 1], [-1, 1]]
    print("Is Reflected (points1):", sol.isReflected(points1))  # Expected output: True
    
    # Example 2:
    points2 = [[1, 1], [2, 2]]
    print("Is Reflected (points2):", sol.isReflected(points2))  # Expected output: False
    
    # Example 3: More points forming symmetry about x=0
    points3 = [[1, 1], [2, 1], [-1, 1], [-2, 1]]
    print("Is Reflected (points3):", sol.isReflected(points3))  # Expected output: True
