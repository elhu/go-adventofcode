package stringset

// StringSet implements common methods for a set of strings
type StringSet struct {
	data map[string]struct{}
}

// New creates an empty string set
func New() *StringSet {
	return &StringSet{data: make(map[string]struct{})}
}

// NewFromSlice creates a string set from a given string slice
func NewFromSlice(s []string) *StringSet {
	set := &StringSet{data: make(map[string]struct{})}
	for _, b := range s {
		set.Add(b)
	}
	return set
}

// Len returns the length of the string set
func (s *StringSet) Len() int {
	return len(s.data)
}

// Members returns the list of members in the set as a slice
func (s *StringSet) Members() []string {
	res := make([]string, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}

// Add adds b to the string set
func (s *StringSet) Add(b string) {
	s.data[b] = struct{}{}
}

// Remove removes b from the string set
func (s *StringSet) Remove(b string) {
	delete(s.data, b)
}

// HasMember returns true if b is present in the string set, false otherwise
func (s *StringSet) HasMember(b string) bool {
	_, found := s.data[b]
	return found
}

// Intersection returns a new string set with the common elements from in the string set and o
func (s *StringSet) Intersection(o *StringSet) *StringSet {
	res := New()
	for k := range s.data {
		if o.HasMember(k) {
			res.Add(k)
		}
	}
	return res
}

// Substraction returns a new string set with the elements of s and minus the elements of o
func (s *StringSet) Substract(o *StringSet) *StringSet {
	res := New()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Remove(k)
	}
	return res
}

// Union returns a new string set with elements of both in the string set and o
func (s *StringSet) Union(o *StringSet) *StringSet {
	res := New()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Add(k)
	}
	return res
}

// Each iterates over each element in the string set and calls fn
func (s *StringSet) Each(fn func(string)) {
	for k := range s.data {
		fn(k)
	}
}
