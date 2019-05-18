package core

// Set type from this gist: https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
type Set struct {
	list map[int]struct{} //empty structs occupy 0 memory
}

// Has the value in the set
func (s *Set) Has(v int) bool {
	_, ok := s.list[v]
	return ok
}

// Add the value in the set
func (s *Set) Add(v int) {
	s.list[v] = struct{}{}
}

// Remove the value in the set
func (s *Set) Remove(v int) {
	delete(s.list, v)
}

// Clear the value in the set
func (s *Set) Clear() {
	s.list = make(map[int]struct{})
}

// Size returns the size of the set
func (s *Set) Size() int {
	return len(s.list)
}

// NewSet creates an empty set
func NewSet() *Set {
	s := &Set{}
	s.list = make(map[int]struct{})
	return s
}

// NewSetFromInts creates a set from ints
func NewSetFromInts(list []int) *Set {
	s := &Set{}
	s.list = make(map[int]struct{})
	s.AddMulti(list...)
	return s
}

//optional functionalities

//AddMulti Add multiple values in the set
func (s *Set) AddMulti(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

// FilterFunc for filtering
type FilterFunc func(v int) bool

// Filter returns a subset, that contains only the values that satisfies the given predicate P
func (s *Set) Filter(P FilterFunc) *Set {
	res := NewSet()
	for v := range s.list {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Union provides the sum of the two sets
func (s *Set) Union(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

// Intersect provides what's common between the two sets
func (s *Set) Intersect(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from s, that doesn't exists in s2 (param)
func (s *Set) Difference(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

// ToSlice converts the set into a slice
func (s *Set) ToSlice() []int {
	res := []int{}
	for v := range s.list {
		res = append(res, v)
	}
	return res
}
