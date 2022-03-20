package maximum_frequency_stack

import (
	"fmt"
	"testing"
)

func TestFreqStack(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	stack.Push(7)
	stack.Push(5)
	stack.Push(7)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack.group)

	if res := stack.Pop(); res != 5 {
		t.Errorf("expected 5, got %v", res)
	}

	if res := stack.Pop(); res != 7 {
		t.Errorf("expected 7, got %v", res)
	}

	if res := stack.Pop(); res != 5 {
		t.Errorf("expected 5, got %v", res)
	}

	if res := stack.Pop(); res != 4 {
		t.Errorf("expected 4, got %v", res)
	}
}
