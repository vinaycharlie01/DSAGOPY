package main

type Deque struct {
	items []int
}

func (deque *Deque) InsertFront(item int) {
	deque.items = append(deque.items, item)
	for i := len(deque.items) - 1; i > 0; i-- {
		deque.items[i] = deque.items[i-1]
	}
	deque.items[0] = item
}
func (deque *Deque) InsertBack(item int) {
	deque.items = append(deque.items, item)
}
func (deque *Deque) First() int {
	return deque.items[0]
}
func (deque *Deque) RemoveFirst() int {
	returnValue := deque.items[0]
	deque.items = deque.items[1:]
	return returnValue
}
func (deque *Deque) Last() int {
	return deque.items[len(deque.items)-1]
}
func (deque *Deque) RemoveLast() int {
	length := len(deque.items)
	returnValue := deque.items[length-1]
	deque.items = deque.items[:(length - 1)]
	return returnValue
}

func (deque *Deque) Empty() bool {
	return len(deque.items) == 0
}
func main() {

}
