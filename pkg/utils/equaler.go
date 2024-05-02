package utils

import (
	"reflect"
)

// Equaler returns true if a and b are equal.
// Uses reflect.DeepEqual to compare a and b.
func Equaler(a, b any) bool {
	return reflect.DeepEqual(a, b)
}
