class DynamicArray:
    
    def __init__(self, capacity: int):
        self.array = [None for _ in range(capacity)]

    def get(self, i: int) -> int:
        if i < self.getSize():
            return self.array[i]
        return 0

    def set(self, i: int, n: int) -> None:
        if i < self.getSize():
            self.array[i] = n

    def pushback(self, n: int) -> None:
        size = self.getSize()
        if size < self.getCapacity():
            self.array[size] = n
        else:
            self.resize()
            self.array[size]= n


    def popback(self) -> int:
        size = self.getSize()
        if size > 0:
            a = self.get(size - 1)
            self.array[size - 1] = None
            return a

    def resize(self) -> None:
        newArray = [None for _ in range(self.getCapacity() * 2)]
        for i in range(self.getSize()):
            newArray[i] = self.array[i]
        self.array = newArray

    def getSize(self) -> int:
        size = 0
        for i in range(len(self.array)):
            if self.array[i] != None:
                size += 1
        return size
    
    def getCapacity(self) -> int:
        return len(self.array)