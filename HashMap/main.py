class Hashmap:
    def __init__(self,data):
        self.data=data
    def chekVal(self,target):
        chek={}
        for i in range(len(self.data)):
            chek[i]=self.data[i]
            if target in chek.values():
                return [True,chek[i],i]
            continue
        return [False,target,-1]
        # /print(chek)
a=Hashmap([1,2,3,4,5])
b=a.chekVal(0)
print(b)

            

    



