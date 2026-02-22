class LRUCache:
    class Node:
        def __init__(self, key, val):
            self.key = key
            self.val = val
            self.prev = None
            self.next = None

    def __init__(self, capacity: int):
        self.cap = capacity
        self.head = self.Node(-1, -1)
        self.tail = self.Node(-1, -1)
        self.head.next = self.tail
        self.tail.prev = self.head
        self.map = {}
    
    def addNode(self, newnode):
        temp = self.head.next
        newnode.next = temp
        newnode.prev = self.head
        self.head.next = newnode
        temp.prev = newnode

    def deleteNode(self, delnode):
        prev = delnode.prev
        next = delnode.next
        prev.next = next
        next.prev = prev

    def get(self, key: int) -> int:
        if key in self.map:
            resNode = self.map[key]
            ans = resNode.val
            del self.map[key]
            self.deleteNode(resNode)
            self.addNode(resNode)
            self.map[key] = self.head.next
            return ans

        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.map:
            cur = self.map[key]
            del self.map[key]
            self.deleteNode(cur)
        
        if len(self.map) == self.cap:
            del self.map[self.tail.prev.key]
            self.deleteNode(self.tail.prev)
        
        self.addNode(self.Node(key, value))
        self.map[key] = self.head.next

