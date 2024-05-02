package hashmap

import "github.com/pietroagazzi/gohashlib/pkg/utils"

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
