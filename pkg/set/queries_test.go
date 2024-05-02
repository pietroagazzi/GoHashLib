package set_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/set"
	"testing"
)

func TestSet_Any(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)

	if !s.Any(func(i int) bool { return i%2 == 0 }) {
		t.Errorf("Expected Any to return true when condition is met")
	}
}

func TestSet_Any_FalseCondition(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)

	if s.Any(func(i int) bool { return i > 3 }) {
		t.Errorf("Expected Any to return false when condition is not met")
	}
}

func TestSet_All(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)

	if !s.All(func(i int) bool { return i > 0 }) {
		t.Errorf("Expected All to return true when condition is met for all elements")
	}
}

func TestSet_All_FalseCondition(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)

	if s.All(func(i int) bool { return i%2 == 0 }) {
		t.Errorf("Expected All to return false when condition is not met for all elements")
	}
}

func TestSet_Contains(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[any](2, 1)
	s2.Add(1, s1, 3)

	if !s2.Contains(s1) {
		t.Errorf("Expected Contains to return true for set")
	}
}
