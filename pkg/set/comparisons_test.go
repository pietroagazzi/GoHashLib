package set_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/set"
	"testing"
)

func TestSet_Equal(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(2, 3, 1)

	if !s1.Equal(s2) {
		t.Errorf("Expected Equal to return true for equal sets")
	}
}

func TestSet_Equal_DifferentLength(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(2, 3)

	if s1.Equal(s2) {
		t.Errorf("Expected Equal to return false for different length sets")
	}
}

func TestSet_Equal_DifferentValues(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(2, 3, 4)

	if s1.Equal(s2) {
		t.Errorf("Expected Equal to return false for different value sets")
	}
}

func TestSet_Equal_NotComparable(t *testing.T) {
	s1 := set.NewSet[any](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[any](2, 1)
	s2.Add(2, 3, s1)

	s3 := set.NewSet[any](4, .3)
	s3.Add(s1, 3, 2)

	if !s2.Equal(s3) {
		t.Errorf("Expected Equal to return true for equal sets")
	}
}

func TestSet_Subset(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2)

	s2 := set.NewSet[int](2, 1)
	s2.Add(2, 1, 3)

	if !s1.Subset(s2) {
		t.Errorf("Expected Subset to return true for subset")
	}

	if s2.Subset(s1) {
		t.Errorf("Expected Subset to return false for non-subset")
	}
}

func TestSet_Subset_False(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(2, 1)

	if s1.Subset(s2) {
		t.Errorf("Expected Subset to return false for non-subset")
	}
}
