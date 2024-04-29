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

// Contains checks if the set contains a value.
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m.Get(value)
	return ok
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(value T) {
	s.m.Delete(value)
}

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

// ToSlice returns a slice containing all values in the set.
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Len())

	for value := range s.m.Iter() {
		slice = append(slice, value.Key)
	}

	return slice
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
