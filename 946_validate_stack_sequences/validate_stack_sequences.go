package validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	p1 := 0
	p2 := 0

	stack := make([]int, 0)

	for {
		if len(stack) > 0 && p2 < len(popped) && stack[len(stack)-1] == popped[p2] {
			stack = stack[:len(stack)-1]
			p2++
		} else if p1 < len(pushed) {
			stack = append(stack, pushed[p1])
			p1++
		} else {
			break
		}
	}

	return p1 == len(pushed) && p2 == len(popped)
}
