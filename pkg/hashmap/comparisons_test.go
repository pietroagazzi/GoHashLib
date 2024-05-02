package hashmap_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/hashmap"
	"testing"
)

func TestMap_Equal(t *testing.T) {
	m1 := hashmap.NewMap[int, string](2, 3)
	m1.Set(1, "one")
	m1.Set(2, "two")

	m2 := hashmap.NewMap[int, string](2, 3)
	m2.Set(1, "one")
	m2.Set(2, "two")

	if !m1.Equal(m2) {
		t.Errorf("Expected maps to be equal")
	}
}

func TestMap_Equal_DifferentValues(t *testing.T) {
	m1 := hashmap.NewMap[int, string](2, 3)
	m1.Set(1, "one")
	m1.Set(2, "two")

	m2 := hashmap.NewMap[int, string](2, 3)
	m2.Set(1, "one")
	m2.Set(2, "three")

	if m1.Equal(m2) {
		t.Errorf("Expected maps to be different")
	}
}

func TestMap_Equal_DifferentLength(t *testing.T) {
	m1 := hashmap.NewMap[int, string](2, 3)
	m1.Set(1, "one")
	m1.Set(2, "two")

	m2 := hashmap.NewMap[int, string](2, 3)
	m2.Set(1, "one")

	if m1.Equal(m2) {
		t.Errorf("Expected maps to be different")
	}
}
