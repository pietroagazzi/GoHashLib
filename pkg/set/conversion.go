package set

// ToSlice returns a slice containing all values in the set.
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Len())

	for value := range s.m.Iter() {
		slice = append(slice, value.Key)
	}

	return slice
}
