class DynamicArray:

    def __init__(self, capacity: int):
        self.array = [None for _ in range(capacity)]
        self.size = 0

    def get(self, i: int) -> int:
        return self.array[i]

    def set(self, i: int, n: int) -> None:
        self.array[i] = n

    def pushback(self, n: int) -> None:
        if self.getSize() == self.getCapacity():
            self.resize()
        self.array[self.getSize()] = n
        self.size += 1

    def popback(self) -> int:
        if self.getSize() > 0:
            element = self.array[self.getSize() - 1]
            self.array[self.getSize() - 1] = None
            self.size -= 1
            return element

    def resize(self) -> None:
        newArray = [None for _ in range(self.getCapacity() * 2)]
        for i in range(self.getSize()):
            newArray[i] = self.array[i]
        self.array = newArray

    def getSize(self) -> int:
        return self.size
    
    def getCapacity(self) -> int:
        return len(self.array)
