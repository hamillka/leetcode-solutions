class RandomizedSet:

    def __init__(self):
        self.lst = []
        self.map = {}

    def insert(self, val: int) -> bool:
        if val in self.map:
            return False
        
        self.lst.append(val)
        self.map[val] = len(self.lst) - 1
        return True

    def remove(self, val: int) -> bool:
        if val not in self.map:
            return False
        
        # [1, 2, 3, 4, 5]
        # {1: 0, 2:1, 3:2, 4:3, 5:4}
        idx = self.map[val] # del 3 //idx=2
        self.lst[idx] = self.lst[-1] # [1, 2, 5, 4, 5]
        self.map[self.lst[-1]] = idx # {1:0, 2:1, 3:2, 4:3, 5:2}
        self.lst.pop() # [1, 2, 5, 4]
        del self.map[val] # {1:0, 2:1, 4:3, 5:2}

        return True

    def getRandom(self) -> int:
        return random.choice(self.lst)

