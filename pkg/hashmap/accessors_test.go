package hashmap_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/hashmap"
	"testing"
)

func TestMap_Keys(t *testing.T) {
	m := hashmap.NewMap[int, string](2, 3)
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")
	m.Set(4, "four")

	keys := m.Keys()

	if len(keys) != 4 {
		t.Errorf("Expected length to be 4, got %d", len(keys))
	}
}

func TestMap_Values(t *testing.T) {
	m := hashmap.NewMap[int, string](2, 3)
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")
	m.Set(4, "four")

	values := m.Values()

	if len(values) != 4 {
		t.Errorf("Expected length to be 4, got %d", len(values))
	}
}
