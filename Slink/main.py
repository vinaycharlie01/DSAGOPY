class Node:
    def __init__(self,data):
        self.data=data
        self.next=None
class Slink:
    def __init__(self):
        self.head=None
        self.tail=None
    
    def append(self,v):
        newNode=Node(v)
        if self.head == None:
            self.head = newNode
        else:
            self.tail.next = newNode
        self.tail = newNode

    def push(self,n,pos):
        if pos == 0:
            newNode =Node(n)
            newNode.next = self.head
            self.head = newNode
        elif  pos > 0 and pos < self.len():
            newNode =Node( n)
            p = self.head
            i = 0
            while i < pos-1:
                p = p.next
                i += 1
            newNode.next = p.next
            p.next = newNode
        else:
            newNode =Node(n)
            self.tail.next = newNode
            self.tail = newNode
    def len(self):
        p = self.head
        count = 0
        while p != None:
            p = p.next
            count+=1
        return count
    def display(self):
        p=self.head
        while p:
            print(p.data,end=" ")
            p=p.next
l = Slink()
l.append(20)
l.append(10)
l.append(2)
l.append(3)
l.append(4)
l.append(60)
l.display()