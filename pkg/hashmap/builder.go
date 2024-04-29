package hashmap

// Entry mocks an entry in the Map.
type Entry[K, V any] struct {
	Key   K
	Value V
}

// Builder is a helper to build a Map.
type Builder[K, V any] []*Entry[K, V]

// Build returns the Map with all the entries.
func (ht *Builder[K, V]) Build(threshold float32) *Map[K, V] {
	m := NewMap[K, V](uint32(len(*ht)), threshold)

	// Add all entries to the Map
	for _, e := range *ht {
		m.Set(e.Key, e.Value)
	}

	return m
}
