
---
# DSA In Go ,Python,C
---

Singley Linked list:







    [head]                                [tail]
      |--[d:a]-->[d:a]-->[d:a]-->[d:a]-->[]--|










Circular Singley  linked list:








    [head]                   
      |--[d:a]-->[d:a]-->[d:a]-->[d:a]-->[d:a]--|
      +-----------------------------------------+  






Doubly Linked list:








     [head]                ---                               [tail]
      |--[a:d:a]<-->[a:d:a]<-->[a:d:a]<-->[a:d:a]<-->[a:d:a]--|







Doubly Circular Linked list:






     [head]                ---                               [tail]
      |--[a:d:a]<-->[a:d:a]<-->[a:d:a]<-->[a:d:a]<-->[a:d:a]--|
      +-------------------------------------------------------+ 







Stack:









    |  10   |----->top
    |-------|           [Last input First output :LIPO]
    |  20   |       
    |-------|      lenth=5
    |  30   |
    |-------|      if this stack is empty then Isemty is true
    |  40   |
    |-------|
    |  50   |---->bottom   [First input Last output :FILO]
    +-------+


Queue:





    FIFO:                                   LILO:


    ----------------------------------------------
    |  50  |  30 | 38  | 34  | 29 | 39 | 93 | 9  |
    ----------------------------------------------

```go

package main

import "fmt"

func main() {
	fmt.Println("Hi Hello....! ")
}

```

---
# Bubblesort
---

```go
    --->its buble up at the last 
    arr =[4,3,2,1]
    lenth of arr=4
    pass1:
         4>3--true
         swap
         [3,4,2,1]
         4>2--true
         swap
         [3,2,4,1]
         4>1--true
         [3,2,1,4]
         swap
        [3,2,1|4]----->bubble up last element
       unsoted|sorted
    pass2:
        [3,2,1]
        3>2---true
        swap
        [2,3,1]
        3>1---true
        swap
        [2,1|3,4]
     unsoted|sorted
    pass3:
        [1,2]
        1>2 ---false
        no swap
    loop will run len(arr)-1 ex:4-1=3
    every time last number bubble so we need to reduce the lenth 
func BabbleSort(arr []int) {
    //pass3 so loop will run lenth-1
	for i := 0; i < len(arr)-1; i++ {
        //value a is bool- if no swap dont require to run loop
		var a bool
        //last value atomatically bubble so lenth-1 and last element sorted so I need to reduce i times
		for j := 0; j < len(arr)-1-i; j++ {
            //if arr[j] element greater then arr[j+1] element
           // if condition is true swap else dont swap
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				a = true
			}
		}
		if !a {
			break
		}
	}

}     

```
---
# SelectionSort
---
```go
  
    --->assume 1st element is min element find the least min element then swap
    arr =[4,3,2,1]
    lenth of arr=4
    pass1:
        min:4
        4<3 --true
           min=3
        2<3 --true
           min=2
        1<2--true
           min=1
        min=1
        swap leastmin_element and 1st element
        [1|3,2,4]
     soted|unsorted
    pass2:
        [3,2,4]
        min=3
        3>2---true
           min=2
        2>4---false
            min=4
        min=2
        swap leastmin_element and 2st element
        [2|3,4]
     soted|unsorted
    pass3:
        [3,4]
        min=3
        3>4---flase
           min=4
        min=3
        no swap
        [2,3,4]
         soted  
    lenth of array =  4
    1st loop will run lenthtimes
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}
```
---
# InsertionSort
---

```go
   one element is sorted and other elements is unsorted
   arrr=[4,3,2,1]
   pass1:
        [4|3,2,1]
     soted| unsorted
        find the corect position of 3 and insert the element 
        [?,4,2,1]
        is there any element left side chek else insert
        [3,4,2,1]
    same method apply pass1
    pass2:
        [3,4|2,1]
       soted| unsorted  
       find the corect position of 2 and insert the element
       [3,?,4,1]
       [?,3,4,1]
       is there any element left side chek else insert
       [2,3,4,1]
    pass3:
        [2,3,4|1]
         soted| unsorted
        find the corect position of 2 and insert the element
        [2,3,?,4]
        [2,?,3,4]
        [?,2,3,4]
        [1,2,3,4]
func InsertionSort(arr []int) {
	//one element is alredy sorted so we run loop index 1 to lenth of array
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}     
```
---
# ShellSort
---

```go
   arr=[9,2,4,1,3,5,6,7,8]
   lenth of arr=9
       find the gap
            gap=len(arr)/2
            start gap=4 and stop gap>0 and  step gap=gap/2
    gap=4
        [9,2,4,1,3,5,6,7,8]
         i       j
        if i>j{
            swap(i,j)
        }else{
            break
        }
        swap                         
        [3,2,4,1,9,5,6,7,8]
             i       j
        no swap
        [3,2,4,1,9,5,6,7,8]
             i       j
        no swap
        [3,2,4,1,9,5,6,7,8]
               i       j
        swap
         [3,2,4,1,9,5,6,7,8]
                  i       j
         [3,2,4,1,8,5,6,7,9]
         chek the back element also
         3>8---false
    gap=2 
        [3,2,4,1,8,5,6,7,9]
         i   j
         no swap
        [3,2,4,1,8,5,6,7,9]
         i   j
        swap
        [3,2,4,1,8,5,6,7,9]
           i   j
        [3,1,4,2,8,5,6,7,9]
           back element -1 index so  dont do anything
        no swap
        [3,1,4,2,8,5,6,7,9]
             i   j
        no swap
        [3,1,4,2,8,5,6,7,9]
               i   j
        swap
        [3,1,4,2,8,5,6,7,9]
                 i   j
        [3,1,4,2,6,5,8,7,9]
             J   i
             no swap
        no swap
        [3,1,4,2,6,5,8,7,9]
                   i   j
        no swap
        [3,1,4,2,6,5,8,7,9]
                     i   j
    gap=1 
        [3,1,4,2,6,5,8,7,9] 
         i j
        swap
        [1,3,4,2,6,5,8,7,9] 
           i j
        no swap
        [1,3,4,2,6,5,8,7,9] 
           i j
        swap
        [1,3,4,2,6,5,8,7,9] 
             j i
        [1,3,2,4,6,5,8,7,9]
           j i
        [1,2,3,4,6,5,8,7,9]
         j i
        
        no swap
        [1,2,3,4,6,5,8,7,9]
               i j
        no swap
        [1,2,3,4,6,5,8,7,9]
                 i j
        no swap
        [1,2,3,4,5,6,8,7,9]
               j i


        no swap
        [1,2,3,4,5,6,8,7,9]
                   i j
        swap
        [1,2,3,4,5,6,8,7,9]
                     i j
        no swap
        [1,2,3,4,5,6,7,8,9]
                   j i
        
         no swap
        [1,2,3,4,5,6,7,8,9]
                       i j 


          
func ShellSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		for j := gap; j < len(arr); j++ {
			for i := j - gap; i >= 0; i = i - gap {
				if arr[i+gap] > arr[i] {
					break
				} else {
					arr[i], arr[i+gap] = arr[i+gap], arr[i]
				}

			}
		}
	}
}
```
---
# Quiksort
---
```go
    quikSort
    arr=[9,2,4,1,10,11,12,8]
    
    [9,2,4,1,10,11,12,8]
    pivot=arr[start]
    start=0
    end=lenth(arr)-1
[pivot element=9]
     | 
    [9,2,4,1,10,11,12,8]
     |               |
    [start]          [end]
    if start ele <=pivot element move start
    if end element is  greate then pivot element move end
[pivot element=9]
     | 
    [9,2,4,1,10,11,12,8]
       |              |
     [start]       [end]
[pivot element=9]
     | 
    [9,2,4,1,10,11,12,8]
         |            |
       [start]      [end]
[pivot element=9]
     | 
    [9,2,4,1,10,11,12,8]
           |          |
         [start]    [end]
[pivot element=9]
     | 
    [9,2,4,1,10,11,12,8]
              |       |
           [start]  [end]
[pivot element=9]
     | 
    [9,2,4,1,8,11,12,10]
                |     |
            [start]  [end]
[pivot element=9]
     | 
    [9,2,4,1,8,11,12,10]
               |   |
           [start][end]
[pivot element=9]
     | 
    [9,2,4,1,8,11,12,10]
               |
           [start]
             [end]
[pivot element=9]
     | 
    [9,2,4,1,8,11,12,10]
             |   |
             | [start]
           [end]
if end<start:
swap(pivot ,end)
[pivot element=9]
     | 
    [8,2,4,1,9,11,12,10]
             |   |
             | [start]
           [end]
   [8,2,4,1][9][11,12,10]
             |
          [sorted]
repeat same method 
   [1,2,4],[8]
    [1,2,4]
func Partion(arr []int, lb int, ub int) int {
	Pivot := arr[lb]
	start := lb
	end := ub
	for start < end {
		if arr[start] <= Pivot {
			start++
		}
		for arr[end] > Pivot {
			end--
		}
		if start < end {
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
		}
	}
	temp := arr[lb]
	arr[lb] = arr[end]
	arr[end] = temp
	return end
}

func QueKSort(arr []int, lb int, ub int) {
	if lb < ub {
		Mid := Partion(arr, lb, ub)
		QueKSort(arr, lb, Mid-1)
		QueKSort(arr, Mid+1, ub)
	}
}

```
---
# Merge sort
---
```go
   arr=[1,2,3,9,6,5,8,7]
   lenth of array =8
   split all the elements
          [1,2,3,9,6,5,8,7]
              /        \
        [ 1,2,3,9]   [6,5,8,7]
           /   \        /  \
       [1,2]  [3,9]  [6,5][8,7]
        /\     /\      /\   /\
      [1][2] [3][9]  [6][5][8][7]
       i  j   i  j    i  j  i  j
       \ /    \ /      \/   \ /
       [1,2]  [3,9]   [5,6] [7,8]  
         i      j       i     j 
         \     /         \    /            
        [1,2,3,9]       [5,6,7,8]
           i               j
            \              /
 output  = [1,2,3,5,6,7,8,9]
func Merge(a []int, b []int) []int {
	L := 0
	R := 0
	arr := []int{}
	for L < len(a) && R < len(b) {
		if a[L] < b[R] {
			arr = append(arr, a[L])
			L++
		} else {
			arr = append(arr, b[R])
			R++
		}
	}
	if L < len(arr) {
		arr = append(arr, a[L:]...)
	}
	if R < len(arr) {
		arr = append(arr, b[R:]...)
	}
	return arr

}

func MergeSort(arr []int) (arr2 []int) {
	if len(arr) < 2 {
		return arr
	}
	a := MergeSort(arr[:len(arr)/2])
	b := MergeSort(arr[len(arr)/2:])
	s := Merge(a, b)
	for i, v := range s {
		arr[i] = v
	}
	return arr

}
   
```