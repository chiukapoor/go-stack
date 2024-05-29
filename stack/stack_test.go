package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := &Stack{}

	stack.Push(1)
	if len(stack.Nums) != 1 {
		t.Errorf("expected stack size 1, got %d", len(stack.Nums))
	}

	stack.Push(2)
	if len(stack.Nums) != 2 {
		t.Errorf("expected stack size 2, got %d", len(stack.Nums))
	}
}

func TestPop(t *testing.T) {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)

	item, err := stack.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if item != 2 {
		t.Errorf("expected 2, got %d", item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if item != 1 {
		t.Errorf("expected 1, got %d", item)
	}

	_, err = stack.Pop()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := &Stack{}

	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("expected stack to be non-empty")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty")
	}
}
