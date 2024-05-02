package set

// Iter returns a channel that yields each value in the set.
func (s *Set[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for value := range s.m.Iter() {
			ch <- value.Key
		}
		close(ch)
	}()
	return ch
}
