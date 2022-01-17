package kth_symbol_in_grammar

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	prev := kthGrammar(n-1, k/2+k%2)

	if k%2 == 0 {
		return prev ^ 1
	} else {
		return prev
	}
}
