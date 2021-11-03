package byte_set

type ByteSet struct {
	data map[byte]struct{}
}

func New() *ByteSet {
	return &ByteSet{data: make(map[byte]struct{})}
}

func (s *ByteSet) Add(b byte) {
	s.data[b] = struct{}{}
}

func (s *ByteSet) Remove(b byte) {
	delete(s.data, b)
}

func (s *ByteSet) IsMember(b byte) bool {
	_, found := s.data[b]
	return found
}

func (s *ByteSet) Intersection(o *ByteSet) *ByteSet {
	res := New()
	for k := range s.data {
		if o.IsMember(k) {
			res.Add(k)
		}
	}
	return res
}

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

func (s *ByteSet) Each(fn func(byte)) {
	for k := range s.data {
		fn(k)
	}
}
