package intset

// IntSet implements common methods for a set of integers
type IntSet struct {
	data map[int]struct{}
}

// New creates an empty integer set
func New() *IntSet {
	return &IntSet{data: make(map[int]struct{})}
}

// NewFromSlice creates an int set from a given int slice
func NewFromSlice(s []int) *IntSet {
	set := &IntSet{data: make(map[int]struct{})}
	for _, b := range s {
		set.Add(b)
	}
	return set
}

// Len returns the length of the integer set
func (s *IntSet) Len() int {
	return len(s.data)
}

// Members returns the list of members in the set as a slice
func (s *IntSet) Members() []int {
	res := make([]int, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}

// Add adds b to the integer set
func (s *IntSet) Add(b int) {
	s.data[b] = struct{}{}
}

// Remove removes b from the integer set
func (s *IntSet) Remove(b int) {
	delete(s.data, b)
}

// HasMember returns true if b is present in the integer set, false otherwise
func (s *IntSet) HasMember(b int) bool {
	_, found := s.data[b]
	return found
}

// Intersection returns a new integer set with the common elements from in the integer set and o
func (s *IntSet) Intersection(o *IntSet) *IntSet {
	res := New()
	for k := range s.data {
		if o.HasMember(k) {
			res.Add(k)
		}
	}
	return res
}

// Substraction returns a new int set with the elements of s and minus the elements of o
func (s *IntSet) Substract(o *IntSet) *IntSet {
	res := New()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Remove(k)
	}
	return res
}

// Union returns a new integer set with elements of both in the integer set and o
func (s *IntSet) Union(o *IntSet) *IntSet {
	res := New()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Add(k)
	}
	return res
}

// Each iterates over each element in the integer set and calls fn
func (s *IntSet) Each(fn func(int)) {
	for k := range s.data {
		fn(k)
	}
}
