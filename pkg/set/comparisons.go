package set

// Equal checks if two sets are equal.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}

	for value := range s.m.Iter() {
		if !other.Contains(value.Key) {
			return false
		}
	}

	return true
}

// Subset checks if the set is a subset of another set.
func (s *Set[T]) Subset(other *Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}

	for value := range s.m.Iter() {
		if !other.Contains(value.Key) {
			return false
		}
	}

	return true
}
