
class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None
        self.prev = None

class Deque:
    def __init__(self):
        self.head = None
        self.tail = self.head

    def isEmpty(self) -> bool:
        if self.head:
            return False
        return True

    def append(self, value: int) -> None:
        newNode = ListNode(value)
        if self.tail:
            newNode.prev = self.tail
            self.tail.next = newNode
            self.tail = newNode
        else:
            self.head = newNode
            self.tail = self.head

    def appendleft(self, value: int) -> None:
        newNode = ListNode(value)
        if self.head:
            newNode.next = self.head
            self.head.prev = newNode
            self.head = newNode
        else:
            self.head = newNode
            self.tail = self.head

    def pop(self) -> int:
        if not self.isEmpty():
            popped = self.tail
            self.tail = self.tail.prev
            if self.tail:
                self.tail.next = None
            else:
                self.head = self.tail
            return popped.val
        return -1

    def popleft(self) -> int:
        if not self.isEmpty():
            popped = self.head
            self.head = self.head.next
            if self.head:
                self.head.prev = None
            else:
                self.tail = self.head
            return popped.val
        return -1

