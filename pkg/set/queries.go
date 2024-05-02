package set

// Any checks if any value in the set satisfies the callback.
func (s *Set[T]) Any(callback func(T) bool) bool {
	for value := range s.m.Iter() {
		if callback(value.Key) {
			return true
		}
	}
	return false
}

// All checks if all values in the set satisfy the callback.
func (s *Set[T]) All(callback func(T) bool) bool {
	for value := range s.m.Iter() {
		if !callback(value.Key) {
			return false
		}
	}
	return true
}

// Contains checks if the set contains a value.
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m.Get(value)
	return ok
}
