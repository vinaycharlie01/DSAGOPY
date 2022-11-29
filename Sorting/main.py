class Sorting:
    def __init__(self,array):
        self.array=array
    '''
    Bub

    '''
    def bubblesort(self):
        arr=self.array
        for i in range(0,len(arr)-1):
            flag=0
            for j in range(0,len(arr)-1-i):
                if arr[j]>arr[j+1]:
                    arr[j],arr[j+1]=arr[j+1],arr[j] 
                    flag=1
            if flag==0:
                break
        return self.array
    def SelectionSort(self):
        arr=self.array
        for i in range(0,len(arr)-1):
            min=i
            for j in range(i+1,len(arr)):
                if arr[min]>arr[j]:
                    min=j
            arr[min],arr[i]=arr[i],arr[min]
        return self.array
    def insertionSort(self):
        arr=self.array
        for i in range(1,len(arr)):
            temp=arr[i]
            j=i-1
            while j>=0 and arr[j]>temp:
                arr[j+1]=arr[j]
                j=j-1
            arr[j+1]=temp
        return self.array
    def shellSort(self):
        arr=self.array
        gap=len(arr)//2
        while gap>0:
            j=gap
            while j<len(arr):
                i=j-gap
                while i>=0:
                    if arr[i]>arr[i+gap]:
                        temp=arr[i]
                        arr[i]=arr[i+gap]
                        arr[i+gap]=temp
                    else:
                        break
                    i=-gap
                j+=1
            gap=gap//2
        return self.array
    def partion(self,lb,ub):
        arr=self.array
        pivot=arr[lb]
        start=lb
        end=ub
        while start<end:
            if arr[start]<=pivot:
                start=start+1
            while arr[end]>pivot:
                end=end-1
            if start<end:
                temp=arr[start]
                arr[start]=arr[end]
                arr[end]=temp
        temp=arr[lb]
        arr[lb]=arr[end]
        arr[end]=temp
        return end
    def quikSortPart(self,lb ,ub):
        if lb<ub:
            mid=self.partion(lb,ub)
            self.quikSortPart(lb, mid-1)
            self.quikSortPart(mid+1,ub)
    def quikSort(self):
        self.quikSortPart(0, len(self.array)-1)
        return self.array
    def merge_sort(self):
        if len(self.array) > 1:
            m = len(self.array)//2
            left = self.array[:m]
            right = self.array[m:]
            leftsorter = Sorting(left)
            leftsorter.merge_sort()
            rightsorter = Sorting(right)
            rightsorter.merge_sort()
            i = 0
            j = 0
            k = 0
            while i < len(left) and j < len(right):
                if left[i] < right[j]:
                    self.array[k] = left[i]
                    i += 1
                else:
                    self.array[k] = right[j]
                    j += 1
                k += 1

            while i < len(left):
                self.array[k] = left[i]
                i += 1
                k += 1

            while j < len(right):
                self.array[k] = right[j]
                j += 1
                k += 1
        return self.array

a=Sorting([1,3,2,4,5,0,0,0])
a.bubblesort()
a.SelectionSort()
a.shellSort()
a.quikSort()
a.merge_sort()
print(a.array)