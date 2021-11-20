package byte_set

// ByteSet implements common methods for a set of bytes
type ByteSet struct {
	data map[byte]struct{}
}

// New creates an empty byte set
func New() *ByteSet {
	return &ByteSet{data: make(map[byte]struct{})}
}

// Add adds b to the byteset
func (s *ByteSet) Add(b byte) {
	s.data[b] = struct{}{}
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
