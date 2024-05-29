package main

import (
	"fmt"
	"sync"

	"github.com/chiukapoor/go-stack/stack"
)

func main() {
	stack := &stack.Stack{
		Nums: []int{},
	}

	var wg sync.WaitGroup

	// Start 50 goroutines to push items to the stack
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			stack.Push(i)
			fmt.Printf("Pushed: %d\n", i)
		}(i)
	}

	// Start 25 goroutines to pop items from the stack
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			item, err := stack.Pop()
			if err != nil {
				fmt.Println("Pop error:", err)
			} else {
				fmt.Printf("Popped: %d\n", item)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Final Stack:", stack.Nums)
}
