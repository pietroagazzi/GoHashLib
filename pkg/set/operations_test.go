package set_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/set"
	"testing"
)

func TestSet_Union(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(3, 4, 5)

	union := s1.Union(*s2)
	elements := map[int]bool{1: false, 2: false, 3: false, 4: false, 5: false}

	for item := range union.Iter() {
		if _, ok := elements[item]; !ok {
			t.Errorf("Union returned unexpected element %v", item)
		}
		elements[item] = true
	}

	for item, present := range elements {
		if !present {
			t.Errorf("Union did not return element %v", item)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(3, 4, 5)

	intersection := s1.Intersection(*s2)

	if intersection.Len() != 1 || !intersection.Contains(3) {
		t.Errorf("Intersection did not return the expected element: %v", intersection)
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := set.NewSet[int](2, 1)
	s1.Add(1, 2, 3)

	s2 := set.NewSet[int](2, 1)
	s2.Add(3, 4, 5)

	difference := s1.Difference(*s2)

	if difference.Len() != 2 || !difference.Contains(1) || !difference.Contains(2) {
		t.Errorf("Difference did not return the expected elements: %v", difference)
	}
}
