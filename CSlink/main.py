class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Clink:
    def __init__(self):
        self.tail = None

    def append(self, n):
        newNode = Node(n)
        if self.tail == None:
            self.tail = newNode
            self.tail.next = newNode
        else:
            newNode.next = self.tail.next
            self.tail.next = newNode
            self.tail = newNode

    def push(self, n, pos):
        if pos == 0:
            newNode = Node(n)
            newNode.next = self.tail.next
            self.tail.next = newNode
        elif pos >= self.len():
            newNode = Node(n)
            newNode.next = self.tail.next
            self.tail.next = newNode
            self.tail = newNode
        elif pos > 0 and pos < self.len():
            newNode = Node(n)
            p = self.tail
            i = 0
            while i < pos:
                p = p.next
                i += 1
            newNode.next = p.next
            p.next = newNode
    def pop(self,pos):
        if pos == 0:
            temp = self.tail.next.next
            self.tail.next = temp
            temp = None
        elif pos >= l.len()-1:
            p = self.tail
            while p.next.next != l.tail:
                p = p.next
            temp = self.tail.next
            p.next.next = temp
            self.tail = p.next
            temp = None
        elif pos > 0 and pos < self.len():
            i = 1
            p = self.tail
            while i < pos:
                p = p.next
                i += 1
            temp = p.next.next
            p.next.next = temp.next
            temp.next = None
    def len(self):
        temp = self.tail
        count = 1
        while self.tail != temp.next:
            count += 1
            temp = temp.next
        return count
    def reverse(self):
        def __init__(self):
            self.prev = None
            self.currNode=None
            self.nextNode=None
        self.currNode= self.tail.next
        self.nextNode = self.currNode.next
        while self.currNode != self.tail:
            self.prev =self.currNode
            self.currNode = self.nextNode
            self.nextNode = self.currNode.next
            self.currNode.next = self.prev
        self.nextNode.next = self.tail
        self.tail = self.nextNode
        p = self.tail
        temp = self.tail
        while self.tail != temp.next:
            temp = temp.next
            print(temp.data, end=" ")
        print(p.data)
    def display(self):
        p = self.tail
        temp = self.tail
        while self.tail != temp.next:
            temp = temp.next
            print(temp.data,end=" ")
        print(p.data)

l = Clink()
l.append(20)
l.append(10)
l.append(2)
l.append(3)
l.append(4)
l.append(60)
l.push(100, 1)
l.pop(4)
l.display()
l.reverse()