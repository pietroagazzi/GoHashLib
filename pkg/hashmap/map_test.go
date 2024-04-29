package hashmap_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/hashmap"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := hashmap.NewMap[int, string](10, 0.75)

	if m.Size() != 10 {
		t.Errorf("Expected size to be 10, got %d", m.Size())
	}
	if m.Threshold != 0.75 {
		t.Errorf("Expected threshold to be 0.75, got %f", m.Threshold)
	}
}

func TestMap_Resize(t *testing.T) {
	m := hashmap.NewMap[int, string](2, 1)

	// Resize the map 2 times
	m.Resize()
	m.Resize()

	if m.Size() != 11 {
		t.Errorf("Expected size to be 5, got %d", m.Size())
	}
}

func TestMap_Index(t *testing.T) {
	m := hashmap.NewMap[interface{}, string](10, 0.75)

	_, err := m.Index(func() {})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestMap_LoadFactor(t *testing.T) {
	m := hashmap.NewMap[int, int](10, 100)

	if m.LoadFactor() != 0 {
		t.Errorf("Expected load factor to be 0, got %f", m.LoadFactor())
	}

	for i := 0; i < 5; i++ {
		m.Set(i, i)
	}

	if m.LoadFactor() != 0.5 {
		t.Errorf("Expected load factor to be 0.5, got %f", m.LoadFactor())
	}
}

func TestMap_Set(t *testing.T) {
	m := hashmap.NewMap[int, string](1, 2)

	m.Set(1, "one")
	m.Set(2, "two")

	m.Set(1, "three")

	if m.Len() != 2 {
		t.Errorf("Expected length to be 2, got %d", m.Len())
	}
}

func TestMap_Get(t *testing.T) {
	m := hashmap.NewMap[int, string](2, 0.75)
	m.Set(1, "one")
	value, ok := m.Get(1)
	if !ok || value != "one" {
		t.Errorf("Expected to get 'one', got '%s'", value)
	}

	// Test for a key that does not exist
	_, ok = m.Get(2)
	if ok {
		t.Errorf("Expected to not find key 2, but it was found")
	}
}

func TestMap_Delete(t *testing.T) {
	t.Run("deleting a key", func(t *testing.T) {
		m := hashmap.NewMap[int, string](2, 3)
		m.Set(1, "one")
		m.Set(2, "two")
		m.Delete(1)
		_, ok := m.Get(1)
		if ok {
			t.Errorf("Expected key 1 to be deleted")
		}
	})

	// Test for a key that does not exist
	t.Run("deleting a key that does not exist", func(t *testing.T) {
		m := hashmap.NewMap[int, string](2, 3)
		m.Set(1, "one")
		m.Delete(2)
		_, ok := m.Get(2)
		if ok {
			t.Errorf("Expected key 2 to not be found")
		}
	})

	t.Run("deleting a key in a chain", func(t *testing.T) {
		m := hashmap.NewMap[int, string](2, 3)
		m.Set(1, "one")
		m.Set(2, "two")
		m.Set(3, "three")
		m.Set(4, "four")
		m.Set(5, "five")

		m.Delete(1)
		_, ok := m.Get(1)
		if ok {
			t.Errorf("Expected key 3 to be deleted")
		}
	})
}

func TestMap_Len(t *testing.T) {
	m := hashmap.NewMap[int, string](2, 3)
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")
	m.Set(4, "four")

	if m.Len() != 4 {
		t.Errorf("Expected length to be 2, got %d", m.Len())
	}
}

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

func TestMap_Clear(t *testing.T) {
	m := hashmap.NewMap[int, string](2, 3)
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")
	m.Set(4, "four")

	m.Clear()

	if m.Len() != 0 {
		t.Errorf("Expected length to be 0, got %d", m.Len())
	}
}
