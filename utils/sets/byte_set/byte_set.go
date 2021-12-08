package byte_set

// ByteSet implements common methods for a set of bytes
type ByteSet struct {
	data map[byte]struct{}
}

// New creates an empty byte set
func New() *ByteSet {
	return &ByteSet{data: make(map[byte]struct{})}
}

// NewFromSlice creates a byte set from a given byte slice
func NewFromSlice(s []byte) *ByteSet {
	set := &ByteSet{data: make(map[byte]struct{})}
	for _, b := range s {
		set.Add(b)
	}
	return set
}

// Len returns the length of the byte set
func (s *ByteSet) Len() int {
	return len(s.data)
}

// Members returns the list of members in the set as a slice
func (s *ByteSet) Members() []byte {
	res := make([]byte, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}

// Add adds b to the byteset
func (s *ByteSet) Add(b byte) {
	s.data[b] = struct{}{}
}

// Equals returns whether or not o is equal to s
func (s *ByteSet) Equals(o *ByteSet) bool {
	if s.Len() != o.Len() {
		return false
	}
	for b := range s.data {
		if !o.IsMember(b) {
			return false
		}
	}
	return true
}

// Remove removes b from the byteset
func (s *ByteSet) Remove(b byte) {
	delete(s.data, b)
}

// IsMember returns true if b is present in the byteset, false otherwise
func (s *ByteSet) IsMember(b byte) bool {
	_, found := s.data[b]
	return found
}

// Intersection returns a new byte set with the common elements from in the byte set and o
func (s *ByteSet) Intersection(o *ByteSet) *ByteSet {
	res := New()
	for k := range s.data {
		if o.IsMember(k) {
			res.Add(k)
		}
	}
	return res
}

// Substraction returns a new byte set with the elements of s and minus the elements of o
func (s *ByteSet) Substract(o *ByteSet) *ByteSet {
	res := New()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Remove(k)
	}
	return res
}

// Union returns a new byte set with elements of both in the byte set and o
func (s *ByteSet) Union(o *ByteSet) *ByteSet {
	res := New()
	for k := range s.data {
		res.Add(k)
	}
	for k := range o.data {
		res.Add(k)
	}
	return res
}

// Each iterates over each element in the byte set and calls fn
func (s *ByteSet) Each(fn func(byte)) {
	for k := range s.data {
		fn(k)
	}
}
