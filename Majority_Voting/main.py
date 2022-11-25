class Majority_Voting:
    def __init__(self,data):
        self.data=data
    def Winner(self):
        candidate = -1
        votes = 0
        for i in range(len(self.data)):
            if votes == 0:
                candidate = self.data[i]
                votes = 1
            elif self.data[i] == candidate:
                votes += 1
            else:
                votes -= 1
        return candidate
a=Majority_Voting([1,1,1,1,1,2,2,2,4,5,2,2,2,2,2,2])
print(a.Winner())