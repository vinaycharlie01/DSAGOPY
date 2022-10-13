class stack:
    def __init__(self):
        self.data=[]
    def  push(self,d):
        self.data.append(d)

    def pop(self):
        if len(self.data)==0:
            print("stack is empty")
        else:
            self.data.pop()
    def display(self):
        i=len(self.data)-1
        while i>=0:
            print(self.data[i])
            i-=1
    def top(self):
        if len(self.data)==0:
            print("list is emety")
        else:
            print(self.data[0])
    

b=stack()
b.push(10)
b.push(20)
b.push(30)
b.push(40)
b.push(50)
b.push(60)
b.pop()
b.display()
b.top()
