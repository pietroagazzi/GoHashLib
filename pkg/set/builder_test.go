package set_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/set"
	"testing"
)

func TestBuilder_Add(t *testing.T) {
	b := new(set.Builder[int])
	b.Add(1)

	if len(*b) != 1 {
		t.Errorf("Expected length to be 1, got %d", len(*b))
	}
}

func TestBuilder_Add_MultipleValues(t *testing.T) {
	b := new(set.Builder[int])
	b.Add(1, 2, 3)

	if len(*b) != 3 {
		t.Errorf("Expected length to be 3, got %d", len(*b))
	}
}

func TestBuilder_Build(t *testing.T) {
	b := new(set.Builder[any])
	b.Add(1, 2, 3, 1)

	s := b.Build(0.75)

	if s.Len() != 3 {
		t.Errorf("Expected length to be 3, got %d", s.Len())
	}
}
