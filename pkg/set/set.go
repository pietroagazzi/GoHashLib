package set

import (
	"fmt"
	"github.com/pietroagazzi/gohashlib/pkg/hashmap"
)

type Set[T any] struct {
	m hashmap.Map[T, bool]
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: *hashmap.NewMap[T, bool](16, 0.75),
	}
}

func (s *Set[T]) Add(value T) {
	s.m.Set(value, true)
}

func (s *Set[T]) AddAll(values ...T) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m.Get(value)
	return ok
}

func (s *Set[T]) Remove(value T) {
	s.m.Delete(value)
}

func (s *Set[T]) Any(callback func(T) bool) bool {
	for value := range s.m.Iter() {
		if callback(value.Key) {
			return true
		}
	}
	return false
}

func (s *Set[T]) All(callback func(T) bool) bool {
	for value := range s.m.Iter() {
		if !callback(value.Key) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Size() int {
	return s.m.Size()
}

func (s *Set[T]) Len() int {
	return s.m.Len()
}

func (s *Set[T]) Clear() {
	s.m.Clear()
}

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

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Len())

	for value := range s.m.Iter() {
		slice = append(slice, value.Key)
	}

	return slice
}

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

func (s *Set[T]) Inspector() *hashmap.Map[T, bool] {
	return &s.m
}
