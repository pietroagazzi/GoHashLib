package go_map

import "reflect"

// Equaler returns true if a and b are equal.
// Uses reflect.DeepEqual to compare a and b.
func Equaler(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
