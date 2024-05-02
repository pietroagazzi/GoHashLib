package set

import (
	"fmt"
	"github.com/pietroagazzi/gohashlib/pkg/hashmap"
)

type Set[T any] struct {
	m hashmap.Map[T, bool]
}

// NewSet creates a new set with the given size and threshold.
func NewSet[T any](size uint32, threshold float32) *Set[T] {
	return &Set[T]{
		m: *hashmap.NewMap[T, bool](size, threshold),
	}
}

// Add adds a value to the set.
func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.m.Set(v, true)
	}
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(value T) {
	s.m.Delete(value)
}

// Size returns the size of the set.
func (s *Set[T]) Size() uint32 {
	return s.m.Size()
}

// Len returns the length of the set.
func (s *Set[T]) Len() int {
	return s.m.Len()
}

// Clear removes all values from the set.
func (s *Set[T]) Clear() {
	s.m.Clear()
}

// Copy returns a copy of the set.
func (s *Set[T]) Copy() *Set[T] {
	copy := NewSet[T](s.m.Size(), s.m.Threshold)
	for value := range s.m.Iter() {
		copy.Add(value.Key)
	}
	return copy
}

// String returns a string representation of the set.
func (s *Set[T]) String() string {
	out := "{"

	for value := range s.m.Iter() {
		out += fmt.Sprintf("%v, ", value.Key)
	}

	if len(out) > 1 {
		out = out[:len(out)-2]
	}

	return out + "}"
}
