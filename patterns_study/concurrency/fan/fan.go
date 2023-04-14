package fan

import "sync"

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

func compute(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func merge(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	out := make(chan int)

	for _, ch := range chs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
