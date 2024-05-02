package set

import "github.com/pietroagazzi/gohashlib/pkg/hashmap"

// Union returns a new set with all the elements that are in either set.
func (s *Set[T]) Union(other Set[T]) *Set[T] {
	result := s.Copy()
	for item := range other.Iter() {
		result.Add(item)
	}
	return result
}

// Intersection returns a new set with all the elements that are in both sets.
func (s *Set[T]) Intersection(other Set[T]) *Set[T] {
	result := NewSet[T](s.Size(), hashmap.DefaultThreshold)

	for item := range s.Iter() {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference returns a new set with all the elements that are in the first set but not in the second set.
func (s *Set[T]) Difference(other Set[T]) *Set[T] {
	result := NewSet[T](s.Size(), hashmap.DefaultThreshold)

	for item := range s.Iter() {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
