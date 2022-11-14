class Node:
    def __init__(self,data):
        self.data=data
        self.next=None
        self.prev=None
class link:
    def __init__(self):
        self.head=None
        self.tail=None
    def append(self,v):
        newNode=Node(v)
        if self.head==None:
            self.head=newNode
            self.tail=newNode
        else:
            self.tail.next=newNode
            newNode.prev=self.tail
            self.tail=newNode
    def pushfirst(self,v):
        newNode=Node(v)
        if self.head ==None:
            return "list is empty"
        else:
            newNode.next=self.head
            self.head.prev=newNode
            self.head=newNode
    def pushany(self,v,pos):
        newNode=Node(v)
        temp=self.head
        i=0
        while i<pos:
            temp=temp.next
            i+=1
        newNode.prev=temp
        newNode.next=temp.next
        temp.next=newNode
        newNode.next.prev=newNode
    def len(self):
        temp=self.head
        count=0
        while temp:
            temp=temp.next
            count+=1
        return count
    def push(self,v,pos):
        if pos==0:
            newNode=Node(v)
            if self.head ==None:
                return "list is empty"
            else:
                newNode.next=self.head
                self.head.prev=newNode
                self.head=newNode
        elif pos>0 and pos<self.len():
            newNode=Node(v)
            temp=self.head
            i=1
            while i<pos:
                temp=temp.next
                i+=1
            newNode.prev=temp
            newNode.next=temp.next
            temp.next=newNode
            newNode.next.prev=newNode
        elif pos==pos<self.len():
            newNode=Node(v)
            temp=self.head
            i=0
            while i<pos:
                temp=temp.next
                i+=1
            newNode.prev=temp
            newNode.next=temp.next
            temp.next=newNode
            newNode.next.prev=newNode
        elif pos >=self.len():
            newNode=Node(v)
            if self.head==None:
                self.head=newNode
                self.tail=newNode
            else:
                self.tail.next=newNode
                newNode.prev=self.tail
            self.tail=newNode

    def delectany(self,pos):
        temp=self.head
        i=0
        while i<pos:
            temp=temp.next
            i+=1
        temp.prev.next=temp.next
        temp.next.prev=temp.prev
    def pop(self,pos):
        if pos==0:
            temp=self.head
            self.head=self.head.next
            self.head.prev=None
        if pos>0 and pos<self.len():
            temp=self.head
            i=0
            while i<pos:
                temp=temp.next
                i+=1
            temp.prev.next=temp.next
            temp.next.prev=temp.prev
        if pos==self.len():
            temp=self.tail
            self.tail=self.tail.prev
            self.tail.next=None
        if pos>self.len():
            print("index out of range")
    def dispalyforword(self):
        p=self.head
        while p:
            print(p.data)
            p=p.next
    def dipplayreve(self):
        p=self.tail
        while p:
            print(p.data)
            p=p.prev

l=link()
l.append(10)
l.append(20)
l.append(30)
l.append(40)
l.append(50)
l.append(60)
l.push(5,9)
l.dispalyforword()
# l.pushany(100,1)
# l.push(42,4,"end")
# l.dispalyforword()