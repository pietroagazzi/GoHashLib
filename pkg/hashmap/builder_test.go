package hashmap_test

import (
	"github.com/pietroagazzi/gohashlib/pkg/hashmap"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	var threshold float32 = 0.75

	builder := hashmap.Builder[string, int]{
		{"c", 1},
		{"a", 2},
	}

	m := builder.Build(threshold)

	if m.Size() != 2 {
		t.Errorf("Expected size to be %d, got %d", 2, m.Size())
	}
	if m.Threshold != threshold {
		t.Errorf("Expected threshold to be %f, got %f", threshold, m.Threshold)
	}
	if m.Len() != 2 {
		t.Errorf("Expected length to be 1, got %d", m.Len())
	}
}
