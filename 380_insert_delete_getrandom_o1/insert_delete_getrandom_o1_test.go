package insert_delete_getrandom_o1

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	set := Constructor()

	if res := set.Insert(1); !res {
		t.Errorf("Expected %v, got %v", true, res)
	}

	if res := set.Remove(2); res {
		t.Errorf("Expected %v, got %v", false, res)
	}

	if res := set.Insert(2); !res {
		t.Errorf("Expected %v, got %v", true, res)
	}

	if res := set.GetRandom(); res != 1 && res != 2 {
		t.Errorf("Expected one of %v, got %v", []int{1, 2}, res)
	}

	if res := set.Remove(1); !res {
		t.Errorf("Expected %v, got %v", true, res)
	}

	if res := set.Insert(2); res {
		t.Errorf("Expected %v, got %v", false, res)
	}

	if res := set.GetRandom(); res != 2 {
		t.Errorf("Expected one of %v, got %v", []int{2}, res)
	}
}

func TestRandomizedSet2(t *testing.T) {
	set := Constructor()

	set.Remove(0)
	set.Remove(0)
	set.Insert(0)
	set.GetRandom()
	set.Remove(0)

	if !set.Insert(0) {
		t.Errorf("Expected true, got false")
	}
}

func TestRandomizedSet3(t *testing.T) {
	set := Constructor()

	set.Insert(0)
	set.Insert(1)
	set.Remove(0)
	set.Insert(2)
	set.Remove(1)

	if res := set.GetRandom(); res != 3 {
		t.Errorf("Expected 2, got %v", res)
	}
}
