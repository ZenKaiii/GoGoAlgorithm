package GoFdn

func stackExample() {
	// Create a stack
	stack := make([]int, 0)
	// push
	stack = append(stack, 10)
	// pop
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	println(v)
	// isEmpty
	if len(stack) == 0 {

	}
}
