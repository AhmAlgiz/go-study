package generator

func String2RuneGenerator(seq string) <-chan rune {
	ch := make(chan rune)
	go func() {
		defer close(ch)
		for _, val := range seq {
			ch <- val
		}
	}()
	return ch
}
