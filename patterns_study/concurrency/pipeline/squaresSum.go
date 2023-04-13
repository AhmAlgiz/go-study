package squaresSum

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, val := range nums {
			out <- val
		}
		close(out)
	}()
	return out
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) int {
	var s int
	for val := range in {
		s += val
	}
	return s
}
