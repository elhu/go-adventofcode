package sets

// Set implements common methods for a generic set
type Set[K comparable] struct {
	data map[K]struct{}
}

// New creates an empty set
func New[K comparable]() *Set[K] {
	return &Set[K]{data: make(map[K]struct{})}
}

// NewFromSlice creates a set from a given slice
func NewFromSlice[K comparable](slice []K) *Set[K] {
	s := &Set[K]{data: make(map[K]struct{})}
	for _, b := range slice {
		s.Add(b)
	}
	return s
}

// Len returns the length of the set
func (s *Set[K]) Len() int {
	return len(s.data)
}

// Members returns the list of members in the set as a slice
func (s *Set[K]) Members() []K {
	res := make([]K, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}

// Add adds b to the set
func (s *Set[K]) Add(b K) {
	s.data[b] = struct{}{}
}

// Remove removes b from the set
func (s *Set[K]) Remove(b K) {
	delete(s.data, b)
}

// HasMember returns true if b is present in the set, false otherwise
func (s *Set[K]) HasMember(b K) bool {
	_, found := s.data[b]
	return found
}

// Intersection returns a new set with the common elements from in the set and o
func (s *Set[K]) Intersection(o *Set[K]) *Set[K] {
	res := New[K]()
	for k := range s.data {
		if o.HasMember(k) {
			res.Add(k)
		}
	}
	return res
}

// Substraction returns a new set with the elements of s and minus the elements of o
func (s *Set[K]) Substract(o *Set[K]) *Set[K] {
	res := New[K]()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Remove(k)
	}
	return res
}

// Union returns a new integer set with elements of both in the set and o
func (s *Set[K]) Union(o *Set[K]) *Set[K] {
	res := New[K]()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Add(k)
	}
	return res
}

// Each iterates over each element in the set and calls fn
func (s *Set[K]) Each(fn func(K)) {
	for k := range s.data {
		fn(k)
	}
}
