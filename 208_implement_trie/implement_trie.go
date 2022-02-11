package implement_trie

type Node struct {
	prefix   rune
	word     bool
	children [26]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			prefix:   0,
			word:     false,
			children: [26]*Node{},
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root

	for _, c := range word {
		key := c - 'a'
		if cur.children[key] == nil {
			cur.children[key] = &Node{
				prefix:   c,
				word:     false,
				children: [26]*Node{},
			}
		}

		cur = cur.children[key]
	}

	cur.word = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root

	for _, c := range word {
		key := c - 'a'
		if cur.children[key] == nil {
			return false
		}

		cur = cur.children[key]
	}

	return cur.word
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root

	for _, c := range prefix {
		key := c - 'a'
		if cur.children[key] == nil {
			return false
		}

		cur = cur.children[key]
	}

	return true
}
