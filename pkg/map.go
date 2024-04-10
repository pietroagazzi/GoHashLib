package go_map

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
)

// Entry represents an item in the hash table.
// Use separate chaining to handle hash collisions.
type Entry[K comparable, V interface{}] struct {
	Key   K
	Value V

	// Next is a pointer to the next item in the chain
	Next *Entry[K, V]
}

// Map represents a Map.
// Use the hash collision resolution technique of separate chaining.
// https://en.wikipedia.org/wiki/Hash_table#Separate_chaining
type Map[K comparable, V interface{}] struct {
	// size is the number of slots in the Map
	size int
	// data is a slice of pointers to slices of Items
	data []*Entry[K, V]

	// Threshold is the maximum load factor before resizing the hash table.
	// Must be a value between 0 and 1. Usually set to 0.75.
	Threshold float64
}

// NewMap returns a new Map with the given size and threshold.
func NewMap[K comparable, V interface{}](size int, threshold float64) Map[K, V] {
	return Map[K, V]{
		size:      size,
		data:      make([]*Entry[K, V], size),
		Threshold: threshold,
	}
}

// Index returns the index of the slot in the hash table where the value should be stored.
//
// It uses Marshal to convert the value to a byte slice, then hashes the byte slice using FNV-1a.
// The hash is then modded by the size of the hash table to get the index.
func (ht Map[K, V]) Index(value K) (index uint32, err error) {
	b, err := json.Marshal(value)

	if err != nil {
		return 0, err
	}

	h := fnv.New32a()
	_, err = h.Write(b)

	return h.Sum32() % uint32(ht.size), err
}

// Set adds an item to the Map.
func (ht Map[K, V]) Set(key K, value V) {
	index, _ := ht.Index(key)
	entry := &Entry[K, V]{Key: key, Value: value}

	// If the slot is empty, create a new slice and add the entry
	if ht.data[index] == nil {
		ht.data[index] = entry
		return
	}

	// If the slot is not empty, check if the key already exists
	current := ht.data[index]
	for current != nil {
		if Equaler(current.Key, key) {
			current.Value = value
			return
		}
		current = current.Next
	}
}

// Get returns the value associated with the key.
func (ht Map[K, V]) Get(key K) (value V, ok bool) {
	index, _ := ht.Index(key)
	current := ht.data[index]

	for current != nil {
		if Equaler(current.Key, key) {
			return current.Value, true
		}

		current = current.Next
	}

	return value, false
}

// Delete removes an item from the Map.
func (ht Map[K, V]) Delete(key K) {
	index, _ := ht.Index(key)
	current := ht.data[index]

	if current == nil {
		return
	}

	if Equaler(current.Key, key) {
		ht.data[index] = current.Next
		return
	}

	for current.Next != nil {
		if Equaler(current.Next.Key, key) {
			// Remove the item from the chain
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Len returns the number of items in the Map.
func (ht Map[K, V]) Len() int {
	count := 0

	for _, entry := range ht.data {
		current := entry
		for current != nil {
			count++
			current = current.Next
		}
	}

	return count
}

// Size returns the size of the Map.
func (ht Map[K, V]) Size() int {
	return ht.size
}

// LoadFactor returns the load factor of the Map.
//
// The load factor is: number of items / number of slots
func (ht Map[K, V]) LoadFactor() float64 { return float64(ht.Len()) / float64(ht.size) }

// Iter returns a channel that iterates over all items in the Map.
func (ht Map[K, V]) Iter() <-chan Entry[K, V] {
	ch := make(chan Entry[K, V])

	go func() {
		for _, entry := range ht.data {
			current := entry
			for current != nil {
				ch <- *current
				current = current.Next
			}
		}
		close(ch)
	}()

	return ch
}

// String returns a string representation of the Map.
func (ht Map[K, V]) String() string {
	str := "{"

	for i := range ht.Iter() {
		str += fmt.Sprintf("%v: %v, ", i.Key, i.Value)
	}

	// Remove the trailing comma and space
	if len(str) > 1 {
		str = str[:len(str)-2]
	}

	return str + "}"
}
