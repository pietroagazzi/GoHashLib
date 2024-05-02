package hashmap

// Iter returns a channel that iterates over all items in the Map.
func (ht *Map[K, V]) Iter() <-chan entry[K, V] {
	ch := make(chan entry[K, V])

	go func() {
		for _, entry := range ht.data {
			current := entry
			for current != nil {
				ch <- *current
				current = current.Next
			}
		}
		close(ch)
	}()

	return ch
}
