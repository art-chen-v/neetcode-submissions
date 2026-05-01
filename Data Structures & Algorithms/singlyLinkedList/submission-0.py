from typing import List

class ListNode:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next

class LinkedList:
    
    def __init__(self):
        self.head = None
        self.tail = None
    
    def get(self, index: int) -> int:
        cur = self.head
        i = 0

        while cur:
            if i == index:
                return cur.val
            cur = cur.next
            i += 1
        
        return -1

    def insertHead(self, val: int) -> None:
        if not self.head:
            self.head = ListNode(val)
            self.tail = self.head
        else:
            newNode = ListNode(val)
            newNode.next = self.head
            self.head = newNode

            
    def insertTail(self, val: int) -> None:
        if not self.head:
            self.head = ListNode(val)
            self.tail = self.head
        else:
            self.tail.next = ListNode(val)
            self.tail = self.tail.next

    def remove(self, index: int) -> bool:
        cur = self.head
        i = 0
        
        while cur:
            if index == 0:
                if cur.next:
                    self.head = cur.next
                else:
                    self.head = None
                    self.tail = None
                return True
            if i + 1 == index:
                if cur.next:
                    if cur.next == self.tail:
                        self.tail = cur
                    cur.next = cur.next.next
                    return True
            cur = cur.next
            i += 1
        return False

    def getValues(self) -> List[int]:
        cur = self.head
        array = []

        while cur:
            array.append(cur.val)
            cur = cur.next
        
        return array
