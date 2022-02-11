package implement_trie

import "testing"

func TestTrie(t *testing.T) {
	tr := Constructor()

	tr.Insert("apple")

	if res := tr.Search("apple"); res != true {
		t.Errorf("search(apple) should be true")
	}

	if res := tr.Search("app"); res != false {
		t.Errorf("search(app) should be false")
	}

	if res := tr.StartsWith("app"); res != true {
		t.Errorf("startsWith(app) should be true")
	}

	tr.Insert("app")

	if res := tr.Search("app"); res != true {
		t.Errorf("search(app) should be true")
	}
}

func TestTrie2(t *testing.T) {
	tr := Constructor()

	if res := tr.Search("majom"); res != false {
		t.Errorf("search(majom) should be false")
	}

	tr.Insert("majom")
	tr.Insert("majom")
	tr.Insert("majom")

	if res := tr.Search("majom"); res != true {
		t.Errorf("search(majom) should be true")
	}
}

func TestTrie3(t *testing.T) {
	tr := Constructor()

	if res := tr.StartsWith("m"); res != false {
		t.Errorf("startsWith(m) should be false")
	}
}