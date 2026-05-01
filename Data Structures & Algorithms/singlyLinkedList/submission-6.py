class ListNode():
    def __init__(self, val, next = None):
        self.val = val
        self.next = next


class LinkedList:

    def __init__(self):
        self.head = ListNode(-1)
        self.tail = self.head

    def get(self, index: int) -> int:
        i = 0
        curr = self.head.next
        while curr:
            if i == index:
                return curr.val
            i += 1
            curr = curr.next
        return - 1

    def insertHead(self, val: int) -> None:
        newHead = ListNode(val)
        newHead.next = self.head.next
        self.head.next = newHead
        if not newHead.next:
            self.tail = newHead

    def insertTail(self, val: int) -> None:
        self.tail.next = ListNode(val)
        self.tail = self.tail.next

    def remove(self, index: int) -> bool:
        curr = self.head
        i = 0

        while i < index and curr:
            curr = curr.next
            i += 1

        if curr and curr.next:
            curr.next = curr.next.next
            if not curr.next:
                self.tail = curr
            return True

        return False

    def getValues(self) -> [int]:
        values = []
        curr = self.head.next
        while curr:
            values.append(curr.val)
            curr = curr.next
        return values
        
