package main

import "fmt"

type Stack []*any

func (s *Stack) Push(v any) {
	*s = append(*s, &v)
}

func (s *Stack) Pop() *any {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v

}

func (s *Stack) Peek() *any {
	v := (*s)[len(*s)-1]
	return v

}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Clear() {
	*s = make([]*any, 0)
}

func isValid(s string) bool {
	st := new(Stack)
	brackets := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := range s {
		if bracket, ok := brackets[s[i]]; ok {
			st.Push(bracket)
		} else if st.Size() > 0 && *st.Peek() == s[i] {
			st.Pop()
		} else {
			return false
		}
	}
	return st.Empty()
}

func main() {
	s := "({})"
	fmt.Println(isValid(s))
}
