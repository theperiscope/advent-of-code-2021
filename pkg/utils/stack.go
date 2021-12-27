package utils

type Stack []string

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(s string) {
	*stack = append(*stack, s)
}

func (stack *Stack) Peek() (element string, found bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		index := len(*stack) - 1
		element := (*stack)[index]
		return element, true
	}
}

func (stack *Stack) Pop() (element string, found bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		index := len(*stack) - 1
		element := (*stack)[index]
		*stack = (*stack)[:index]
		return element, true
	}
}
