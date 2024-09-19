class Array :
    def __init__(self) :
        self.array = []
    def append(slef , value) :
        slef.array.append(value)
    def get(self , index) :
        return self.array[index]
    def len(self) :
        return len(self.array)


arr = Array()
arr.append(1)
arr.append(2)
arr.append(3)
arr.append(4)
arr.append(5)

print(arr.get(0))
print(arr.get(1))
print(arr.len())
