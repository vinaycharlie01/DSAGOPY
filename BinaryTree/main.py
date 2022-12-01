class Node:
    def __init__(self, data): 
        self.data = data 
        self.left = None  
        self.right = None 
    # def insert(self, data):
    #     if self.data == data:
    #         return False
    #     elif data < self.data:
    #         if self.left:
    #             return self.left.insert(data)
    #         else:
    #             self.left = Node(data)
    #             return True
    #     else:
    #         if self.right:
    #             return self.right.insert(data)
    #         else:
    #             self.right= Node(data)
    #             return True
class BinarySearchTree:
    def __init__(self): 
        self.root = None
    def create(self, val):  
        # if root is Nil
        if self.root==None:
            self.root = Node(val)
        else:
            current = self.root
            while True:
                if val < current.data:
                    if current.left:
                        current = current.left
                    else:
                        current.left = Node(val)
                        break
                elif val > current.data:
                    if current.right:
                        current = current.right
                    else:
                        current.right = Node(val)
                        break
                else:
                    break


    # using recursive function
    def insert(self, val):
        # if root is nil
        if self.root == None:
            self.root = Node(val)
        else:
            self._insert(self.root,val)
    def _insert(self,root,value):
        # if value <root.value then left child data =val
        if value < root.data:
            if root.left == None:
                root.left = Node(value)
            else:
                self._insert(root.left,value)
        # if value >root.value then left child data =val
        elif value >= root.data:
            if root.right == None:
                # root.right = Node(value)
                root.right = Node(value)
            else:
                self._insert(root.right,value)
    def preOrder(self,root):
        if root == None:
            return
        print(root.data, end=" ")
        self.preOrder(root.left)
        self.preOrder(root.right)
    def inOrder(self,root):
        if root == None:
            return
        self.inOrder(root.left)
        print(root.data, end=" ")
        self.inOrder(root.right)
    def postOrder(self,root):
        if root == None:
            return
        self.postOrder(root.left)
        self.postOrder(root.right)
        print(root.data, end=" ")
   



b=BinarySearchTree()

a=[1,9,2,9,3,7,4,8,2,7]
for i in a:
    b.insert(i)

b.preOrder(b.root)
print()
b.inOrder(b.root)
print()
b.postOrder(b.root)
print()

