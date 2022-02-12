package lru_cache

import "testing"

func TestLRUCache(t *testing.T) {
	c := Constructor(2)

	c.Put(1, 1)
	c.Put(2, 2)

	if res := c.Get(1); res != 1 {
		t.Errorf("Expected %v, got %v", 1, res)
	}

	c.Put(3, 3)

	if res := c.Get(2); res != -1 {
		t.Errorf("Expected %v, got %v", -1, res)
	}

	c.Put(4, 4)

	if res := c.Get(1); res != -1 {
		t.Errorf("Expected %v, got %v", -1, res)
	}

	if res := c.Get(3); res != 3 {
		t.Errorf("Expected %v, got %v", 3, res)
	}

	if res := c.Get(4); res != 4 {
		t.Errorf("Expected %v, got %v", 4, res)
	}
}

func TestLRUCache2(t *testing.T) {
	c := Constructor(1)

	c.Put(1, 5)
	c.Put(1, 6)
	c.Put(1, 7)

	if res := c.Get(1); res != 7 {
		t.Errorf("Expected %v, got %v", 7, res)
	}
}

func TestLRUCache3(t *testing.T) {
	c := Constructor(2)

	c.Put(1, 5)
	c.Put(2, 10)

	if res := c.Get(1); res != 5 {
		t.Errorf("Expected %v, got %v", 5, res)
	}

	c.Put(1, 6)
	c.Put(2, 11)
	c.Put(1, 7)

	if res := c.Get(2); res != 11 {
		t.Errorf("Expected %v, got %v", 11, res)
	}

	c.Put(3, 15)

	if res := c.Get(3); res != 15 {
		t.Errorf("Expected %v, got %v", 15, res)
	}

	if res := c.Get(1); res != -1 {
		t.Errorf("Expected %v, got %v", -1, res)
	}

	if res := c.Get(2); res != 11 {
		t.Errorf("Expected %v, got %v", 11, res)
	}
}

func TestLRUCache4(t *testing.T) {
	c := Constructor(3)

	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)

	c.Get(4)
	c.Get(3)
	c.Get(2)

	if res := c.Get(1); res != -1 {
		t.Errorf("Expected %v, got %v", -1, res)
	}

	c.Put(5, 5)
	c.Get(1)

	if res := c.Get(2); res != 2 {
		t.Errorf("Expected %v, got %v", 2, res)
	}

	c.Get(3)

	if res := c.Get(4); res != -1 {
		t.Errorf("Expected %v, got %v", -1, res)
	}

	c.Get(5)
}
