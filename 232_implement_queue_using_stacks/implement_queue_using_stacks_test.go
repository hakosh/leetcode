package implement_queue_using_stacks

import (
	"testing"
)

func TestMyQueue(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)

	if v := q.Peek(); v != 1 {
		t.Errorf("Expected %v, got %v", 1, v)
	}

	if v := q.Pop(); v != 1 {
		t.Errorf("Expected %v, got %v", 1, v)
	}

	if v := q.Empty(); v {
		t.Errorf("Expected %v, got %v", false, v)
	}

	if v := q.Pop(); v != 2 {
		t.Errorf("Expected %v, got %v", 2, v)
	}

	if v := q.Empty(); !v {
		t.Errorf("Expected %v, got %v", true, v)
	}
}

func TestMyQueue2(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Pop()
	q.Push(5)
	q.Pop()
	q.Pop()
	q.Pop()

	if v := q.Pop(); v != 5 {
		t.Errorf("Expected %v, got %v", 5, v)
	}
}
