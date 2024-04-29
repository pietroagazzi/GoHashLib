package set

type Builder[K any] []K

func (sb *Builder[K]) Add(items ...K) {
	*sb = append(*sb, items...)
}

func (sb *Builder[K]) Build(threshold float32) *Set[K] {
	s := NewSet[K](uint32(len(*sb)), threshold)

	for _, item := range *sb {
		s.Add(item)
	}

	return s
}
