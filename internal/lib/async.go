package lib

func Async[T any](f func() T) <-chan T {
	c := make(chan T, 1)
	go func() {
		defer close(c)
		c <- f()
	}()
	return c
}
