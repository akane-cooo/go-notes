package ch01

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](size int) *Set[T] {
	set := make(Set[T], size)
	return &set
}

func (s Set[T]) Add(ele T) {
	s[ele] = struct{}{}
}

func (s Set[T]) Remove(ele T) {
	delete(s, ele)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Has(ele T) bool {
	_, exists := s[ele]
	return exists
}

func (s Set[T]) Range(fn func(ele T)) {
	for key := range s {
		fn(key)
	}
}
