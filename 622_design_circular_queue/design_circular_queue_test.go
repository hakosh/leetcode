package design_circular_queue

import "testing"

func TestMyCircularQueue(t *testing.T) {
	queue := Constructor(3)

	if v := queue.EnQueue(1); !v {
		t.Errorf("Expected %v, got %v", true, v)
	} // return True

	if v := queue.EnQueue(2); !v {
		t.Errorf("Expected %v, got %v", true, v)
	} // return True

	if v := queue.EnQueue(3); !v {
		t.Errorf("Expected %v, got %v", true, v)
	} // return True

	if v := queue.EnQueue(4); v {
		t.Errorf("Expected %v, got %v", false, v)
	} // return False

	if v := queue.Rear(); v != 3 {
		t.Errorf("Expected %v, got %v", 3, v)
	} // return 3

	if v := queue.IsFull(); !v {
		t.Errorf("Expected %v, got %v", true, v)
	} // return True

	if v := queue.DeQueue(); !v {
		t.Errorf("Expected %v, got %v", true, v)
	} // return True

	if v := queue.EnQueue(4); !v {
		t.Errorf("Expected %v, got %v", true, v)
	} // return True

	if v := queue.Rear(); v != 4 {
		t.Errorf("Expected %v, got %v", 4, v)
	} // return 4
}

func TestMyCircularQueue2(t *testing.T) {
	queue := Constructor(3)

	queue.EnQueue(2)

	if queue.Rear() != queue.Front() {
		t.Errorf("With only one item rear and front should be the same, instead front: %v, rear: %v", queue.Front(), queue.Rear())
	}
}

func TestMyCircularQueue3(t *testing.T) {
	queue := Constructor(81)

	queue.EnQueue(69)
	queue.DeQueue()
	queue.EnQueue(92)
	queue.EnQueue(12)
	queue.DeQueue()

	if queue.Front() != 12 {
		t.Errorf("Front should be 12, instead front: %v", queue.Front())
	}
}
