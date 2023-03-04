package set

func New[T comparable](values ...T) *set[T] {
	s := &set[T]{
		v: make(map[T]bool),
	}
	s.Add(values...)
	return s
}

type set[T comparable] struct {
	v map[T]bool
}

func (s *set[T]) Add(values ...T) {
	for _, value := range values {
		s.v[value] = true
	}
}
func (s *set[T]) Has(values ...T) bool {
	mustHaveAll := 0
	for _, value := range values {
		if _, ok := s.v[value]; ok {
			mustHaveAll++
		}
	}
	return len(values) == mustHaveAll
}

func (s *set[T]) Delete(values ...T) {
	for _, value := range values {
		delete(s.v, value)
	}
}
func (s *set[T]) Values() []T {
	ret := make([]T, 0)
	for k, v := range s.v {
		if v {
			ret = append(ret, k)
		}

	}
	return ret
}
