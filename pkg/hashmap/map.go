package hashmap

import (
	"encoding/json"
	"fmt"
	"github.com/pietroagazzi/gohashlib/pkg/utils"
	"hash/fnv"
)

const DefaultThreshold = 0.75

// Map represents a Map.
// Use the hash collision resolution technique of separate chaining.
// https://en.wikipedia.org/wiki/Hash_table#Separate_chaining
type Map[K, V any] struct {
	// size is the number of slots in the Map
	size uint32
	// data is a slice of pointers to slices of Items
	data []*entry[K, V]

	// Threshold is the maximum load factor before resizing the hash table.
	// Must be a value between zero and one.
	// Usually set to 0.75.
	Threshold float32
}

// NewMap returns a new Map with the given size and threshold.
func NewMap[K, V any](size uint32, threshold float32) *Map[K, V] {
	return &Map[K, V]{
		size:      size,
		data:      make([]*entry[K, V], size),
		Threshold: threshold,
	}
}

// Resize changes the size of the Map.
//
// The new size is calculated by doubling the current size and finding the Next prime number.
// https://planetmath.org/goodhashtableprimes suggests using prime numbers for the size of the hash table.
// This helps reduce collisions and distribute the items more evenly.
func (ht *Map[K, V]) Resize() {
	// Find the Next prime number after doubling the size
	ht.size = uint32(utils.NextPrime(int(ht.size) * 2))
	newData := make([]*entry[K, V], ht.size)

	// If no items in the hash table, return
	if ht.Len() == 0 {
		ht.data = newData
		return
	}

	// Copy and rehash the items
	for entry := range ht.Iter() {
		index, _ := ht.Index(entry.Key)
		entry.Next = newData[index]
		newData[index] = &entry
	}

	ht.data = newData
}

// Index returns the index of the slot in the hash table where the value should be stored.
//
// It uses Marshal to convert the value to a byte slice, then hashes the byte slice using FNV-1a.
// The hash is then modded by the size of the hash table to get the index.
func (ht *Map[K, V]) Index(value K) (index uint32, err error) {
	b, err := json.Marshal(value)

	if err != nil {
		return 0, err
	}

	h := fnv.New32a()
	_, err = h.Write(b)

	return h.Sum32() % ht.size, err
}

// Set adds an item to the Map.
func (ht *Map[K, V]) Set(key K, value V) {
	// If the size is zero, create a new slice
	if ht.size == 0 {
		// Resize to 2 if the size is zero
		ht.Resize()
	}

	index, _ := ht.Index(key)
	newEntry := &entry[K, V]{Key: key, Value: value}

	// If the slot is empty, create a new slice and add the newEntry
	if ht.data[index] == nil {
		ht.data[index] = newEntry

		// If the capacity is reached, resize the hash table
		if ht.LoadFactor() >= ht.Threshold {
			ht.Resize()
		}

		return
	}

	// If the slot is not empty, check if the key already exists
	current := ht.data[index]
	for current != nil {
		// If the key already exists, update the value
		if utils.Equaler(current.Key, key) {
			current.Value = value
			return
		}
		current = current.Next
	}

	// If the key does not exist, add the newEntry to the chain
	newEntry.Next = ht.data[index]
	ht.data[index] = newEntry
}

// Get returns the value associated with the key.
func (ht *Map[K, V]) Get(key K) (value V, ok bool) {
	index, _ := ht.Index(key)
	current := ht.data[index]

	for current != nil {
		if utils.Equaler(current.Key, key) {
			return current.Value, true
		}

		current = current.Next
	}

	return value, false
}

// Delete removes an item from the Map.
func (ht *Map[K, V]) Delete(key K) {
	index, _ := ht.Index(key)
	current := ht.data[index]

	if current == nil {
		return
	}

	if utils.Equaler(current.Key, key) {
		ht.data[index] = current.Next
		return
	}

	// If the item is in the chain, remove it
	for current.Next != nil {
		if utils.Equaler(current.Next.Key, key) {
			// Remove the item from the chain
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Len returns the number of items in the Map.
func (ht *Map[K, V]) Len() int {
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
func (ht *Map[K, V]) Size() uint32 {
	return ht.size
}

// LoadFactor returns the load factor of the Map.
//
// The load factor is: number of items / *number of slots*
func (ht *Map[K, V]) LoadFactor() float32 { return float32(ht.Len()) / float32(ht.size) }

// Iter returns a channel that iterates over all items in the Map.
func (ht *Map[K, V]) Iter() <-chan entry[K, V] {
	ch := make(chan entry[K, V])

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

// Clear removes all items from the Map.
func (ht *Map[K, V]) Clear() {
	ht.data = make([]*entry[K, V], ht.size)
}

// Equal returns true if the Map is equal to another Map.
func (ht *Map[K, V]) Equal(other *Map[K, V]) bool {
	if ht.Len() != other.Len() {
		return false
	}

	for i := range ht.Iter() {
		value, ok := other.Get(i.Key)

		if !ok || !utils.Equaler(value, i.Value) {
			return false
		}
	}

	return true
}

// String returns a string representation of the Map.
func (ht *Map[K, V]) String() string {
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
