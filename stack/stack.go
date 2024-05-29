package stack

import (
	"errors"
	"sync"
)

// Stack is represented as a struct of list, it has three functions: Pop, Push, and IsEmpty
type Stack struct {
	Nums []int
	lock sync.Mutex
}

// Pop removes the top item from the stack and returns it
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return 0, errors.New("empty stack")
	}
	item := s.Nums[len(s.Nums)-1]
	s.Nums = s.Nums[:len(s.Nums)-1]
	return item, nil
}

// Push adds an item to the stack
func (s *Stack) Push(num int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Nums = append(s.Nums, num)
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.Nums) == 0
}
