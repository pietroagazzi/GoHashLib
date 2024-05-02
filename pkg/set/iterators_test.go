package set_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/set"
	"testing"
)

func TestSet_Iter(t *testing.T) {
	s := set.NewSet[int](2, 1)
	s.Add(1, 2, 3)

	iter := s.Iter()

	for i := 1; i <= 3; i++ {
		if !s.Contains(i) {
			t.Errorf("Expected set to contain %d", i)
		}
	}

	for i := 1; i <= 3; i++ {
		if value, ok := <-iter; !ok || !s.Contains(value) {
			t.Errorf("Expected set to contain %d", value)
		}
	}

	if _, ok := <-iter; ok {
		t.Errorf("Expected iterator to be closed")
	}
}
