package design_linked_list

import (
	"testing"
)

func TestMyLinkedList_AddAtHead(t *testing.T) {
	ll := Constructor()

	ll.AddAtHead(18)
	if v := ll.Get(0); v != 18 {
		t.Errorf("Expected %v, got %v\n", 18, v)
	}

	ll.AddAtHead(5)
	if v := ll.Get(0); v != 5 {
		t.Errorf("Expected %v, got %v\n", 5, v)
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	ll := Constructor()

	ll.AddAtTail(18)
	if v := ll.Get(0); v != 18 {
		t.Errorf("Expected %v, got %v\n", 18, v)
	}

	ll.AddAtTail(5)
	if v := ll.Get(1); v != 5 {
		t.Errorf("Expected %v, got %v\n", 5, v)
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	ll := Constructor()

	ll.AddAtHead(1)
	ll.AddAtHead(2)
	ll.AddAtHead(3)

	ll.AddAtIndex(1, 20)
	if v := ll.Get(1); v != 20 {
		t.Errorf("Expected %v, got %v\n", 20, v)
	}

	ll.AddAtIndex(3, 19)
	if v := ll.Get(3); v != 19 {
		t.Errorf("Expected %v, got %v\n", 19, v)
	}

	ll.AddAtIndex(25, 10)
}

func TestMyLinkedList_Get(t *testing.T) {
	ll := Constructor()

	ll.AddAtHead(13)
	ll.AddAtHead(12)
	ll.AddAtHead(11)

	if v := ll.Get(1); v != 12 {
		t.Errorf("Expected %v, got %v\n", 12, v)
	}

	if v := ll.Get(2); v != 13 {
		t.Errorf("Expected %v, got %v\n", 13, v)
	}

	if v := ll.Get(8); v != -1 {
		t.Errorf("Expected %v, got %v\n", -1, v)
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	ll := Constructor()

	ll.AddAtHead(2)
	ll.AddAtTail(3)
	ll.AddAtTail(4)
	ll.AddAtTail(5)

	ll.DeleteAtIndex(2)

	if v := ll.Get(2); v != 5 {
		t.Errorf("Expected %v, got %v\n", 5, v)
	}

	ll.DeleteAtIndex(0)

	if v := ll.Get(0); v != 3 {
		t.Errorf("Expected %v, got %v\n", 3, v)
	}

	ll.DeleteAtIndex(8)
}

func TestMyLinkedList(t *testing.T) {
	ll := Constructor()

	ll.AddAtHead(7)
	ll.AddAtHead(2)
	ll.AddAtHead(1)

	if v := ll.Get(0); v != 1 {
		t.Errorf("Expected %v, got %v\n", 1, v)
	}

	ll.AddAtIndex(3, 0)

	if v := ll.Get(3); v != 0 {
		t.Errorf("Expected %v, got %v\n", 0, v)
	}

	ll.DeleteAtIndex(2)
	ll.AddAtHead(6)
	ll.AddAtTail(4)

	if v := ll.Get(4); v != 4 {
		t.Errorf("Expected %v, got %v\n", 4, v)
	}
}

func TestMyLinkedList_Get2(t *testing.T) {
	ll := Constructor()

	ll.AddAtIndex(0, 10)
	ll.AddAtIndex(0, 20)
	ll.AddAtIndex(1, 30)

	if v := ll.Get(0); v != 20 {
		t.Errorf("Expected %v, got %v\n", 20, v)
	}
}
