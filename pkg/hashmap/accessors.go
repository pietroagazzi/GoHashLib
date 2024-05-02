package hashmap

// Keys return a slice of all keys in the Map.
func (ht *Map[K, V]) Keys() []K {
	keys := make([]K, 0)

	for i := range ht.Iter() {
		keys = append(keys, i.Key)
	}

	return keys
}

// Values return a slice of all values in the Map.
func (ht *Map[K, V]) Values() []V {
	values := make([]V, 0)

	for i := range ht.Iter() {
		values = append(values, i.Value)
	}

	return values
}
