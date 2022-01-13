package implement_stack_using_queues

import "testing"

func TestMyStack(t *testing.T) {
	s := Constructor()

	s.Push(1)
	s.Push(2)

	if v := s.Top(); v != 2 {
		t.Errorf("Expected %v, got %v", 2, v)
	}

	if v := s.Pop(); v != 2 {
		t.Errorf("Expected %v, got %v", 2, v)
	}

	if v := s.Empty(); v {
		t.Errorf("Expected %v, got %v", false, v)
	}

	if v := s.Pop(); v != 1 {
		t.Errorf("Expected %v, got %v", 1, v)
	}

	if v := s.Empty(); !v {
		t.Errorf("Expected %v, got %v", true, v)
	}
}
