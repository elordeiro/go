package set

type Set map[any]struct{}

func NewSet() *Set {
	s := make(Set)
	return &s
}

func (s *Set) Add(item any) *Set {
	if s == nil {
		s := make(Set)
		s[item] = struct{}{}
		return &s
	}
	(*s)[item] = struct{}{}
	return s
}

func (s *Set) Remove(item any) {
	delete(*s, item)
}

func (s *Set) Contains(item any) bool {
	if s == nil {
		return false
	}
	_, found := (*s)[item]
	return found
}
