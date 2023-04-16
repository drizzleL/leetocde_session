package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	var idx int
	for _, v := range popped {
		for (len(stack) == 0 || stack[len(stack)-1] != v) && idx < len(pushed) { // just match
			stack = append(stack, pushed[idx])
			idx++
		}
		if stack[len(stack)-1] != v {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}
