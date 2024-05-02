package hashmap

// entry represents an item in the hash table.
// Use separate chaining to handle hash collisions.
type entry[K, V any] struct {
	Key   K
	Value V

	// Next is a pointer to the Next item in the chain
	Next *entry[K, V]
}
