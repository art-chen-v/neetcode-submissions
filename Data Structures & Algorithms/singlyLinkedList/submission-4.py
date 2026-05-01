class ListNode():
    def __init__(self, val):
        self.val = val
        self.next = None


class LinkedList:

    def __init__(self):
        self.head = None
        self.tail = None

    def get(self, index: int) -> int:
        count = 0
        curr = self.head

        while curr is not None:
            if count == index:
                return curr.val
            count += 1
            curr = curr.next
        return - 1

    def insertHead(self, val: int) -> None:
        newHead = ListNode(val)
        newHead.next = self.head
        if newHead.next is None:
            self.tail = newHead
        self.head = newHead

    def insertTail(self, val: int) -> None:
        if self.head is None:
            self.insertHead(val)
        else:
            self.tail.next = ListNode(val)
            self.tail = self.tail.next

    def remove(self, index: int) -> bool:
        curr = self.head
        count = 1

        if index == 0:
            if self.head is not None:
                self.head = curr.next
                return True
            return False

        while curr.next is not None:
            if index == count:
                curr.next = curr.next.next
                if curr.next is None:
                    self.tail = curr
                return True
            curr = curr.next
            count += 1
        return False

    def getValues(self) -> List[int]:
        values = []
        curr = self.head
        while curr is not None:
            values.append(curr.val)
            curr = curr.next
        return values
