class Node:
    def __init__(self,data):
        self.data=data
        self.next=None
class Stacklink:
    def __init__(self):
        self.top=None
    def push(self,v):
        newNode=Node(v)
        if self.top==None:
            self.top=newNode
        else:
            newNode.next=self.top
            self.top=newNode
    def pop(self):
        if self.top==None:
            print("Under flow")
        else:
            # print(self.top.data)
            temp=self.top.next
            self.top=temp
            temp=None
    def peek(self):
        if self.top==None:
            print("stack is empty")
        else:
            print("TOP element")
            print("[",self.top.data,"  ]")
            print()
    def display(self):
        p=self.top
        while p !=None:
            print(  "|","",p.data,"", "|")
            print("|------|")
            p=p.next
l=Stacklink()
l.push(10)
l.push(20)
l.push(30)
l.push(40)
l.push(50)
print("stack element before pop")
l.display()
l.pop()
l.pop()
print("stack element after two element pop")
l.display()

print("this is top element of stack")
l.peek()