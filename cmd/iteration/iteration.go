package iteration

func Iteration(input string, count int) string {
	var repeat string
	for i := 0; i < count; i++ {
		repeat += input
	}
	return repeat
}
