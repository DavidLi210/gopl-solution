package exe9_4

func pipeline(stages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for num := range in {
				out <- num
			}
			close(out)
		}(in, out)
	}
	return first, out
}
