class Node:
    def __init__(self, val):
        self.val = val
        self.next = None
        self.prev = None

class Deque:
    
    def __init__(self):
        self.left = Node(-1)
        self.right = Node(-1)
        self.left.next = self.right
        self.right.prev = self.left

    def isEmpty(self) -> bool:
        return self.left.next == self.right

    def append(self, value: int) -> None:
        newNode, prevNode = Node(value), self.right.prev
        prevNode.next = newNode
        newNode.prev = prevNode
        newNode.next = self.right
        self.right.prev = newNode

    def appendleft(self, value: int) -> None:
        newNode, nextNode = Node(value), self.left.next
        self.left.next = newNode
        newNode.prev = self.left
        newNode.next = nextNode
        nextNode.prev = newNode

    def pop(self) -> int:
        if self.isEmpty():
            return -1
        target = self.right.prev
        self.right.prev = target.prev
        target.prev.next = self.right
        return target.val

    def popleft(self) -> int:
        if self.isEmpty():
            return -1
        target = self.left.next
        self.left.next = target.next
        target.next.prev = self.left
        return target.val
