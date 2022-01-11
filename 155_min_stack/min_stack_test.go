package min_stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := Constructor()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if v := stack.GetMin(); v != -3 {
		t.Errorf("Expected %v, got %v", -3, v)
	}

	stack.Pop()

	if v := stack.Top(); v != 0 {
		t.Errorf("Expected %v, got %v", 0, v)
	}

	if v := stack.GetMin(); v != -2 {
		t.Errorf("Expected %v, got %v", -2, v)
	}
}
