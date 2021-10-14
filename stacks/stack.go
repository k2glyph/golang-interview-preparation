package main

import "fmt"

// stack to store values
type stack struct {
	items []int
}

// push
func (s *stack) push(item int) {
	s.items = append(s.items, item)
}

// pop
func (s *stack) pop() int {
	itemToRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return itemToRemove
}
func main() {
	myStack := stack{}
	myStack.push(10)
	myStack.push(20)
	myStack.push(30)
	fmt.Println(myStack)
	myStack.pop()
	// myStack.pop()
	fmt.Println(myStack)
}
