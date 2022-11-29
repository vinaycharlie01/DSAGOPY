This is c programming dsa linked list

i will learn Linked List Programming in Go python  and c



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

```go
   Bubblesort --->its buble up at the last 
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
#SelectionSort
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