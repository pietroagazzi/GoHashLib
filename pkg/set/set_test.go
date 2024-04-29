package set_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/set"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1)

	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1 after adding")
	}
}

func TestSet_Add_MultipleValues(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3, 2)

	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("Expected set to contain 1, 2, 3 after adding all")
	}
}

func TestSet_Remove(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1)
	s.Remove(1)

	if s.Contains(1) {
		t.Errorf("Expected set to not contain 1 after removing")
	}
}

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

func TestSet_Clear(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)
	s.Clear()

	if s.Len() != 0 {
		t.Errorf("Expected set to be empty after Clear")
	}
}

func TestSet_ToSlice(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)
	slice := s.ToSlice()

	if len(slice) != 3 || !s.Contains(slice[0]) || !s.Contains(slice[1]) || !s.Contains(slice[2]) {
		t.Errorf("Expected ToSlice to return a slice containing all elements in the set")
	}
}

func TestSet_String(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3, 4, 5)

	s.Remove(3)

	if s.String() != "{5, 4, 1, 2}" {
		t.Errorf("Expected String to return a string representation of the set")
	}
}
