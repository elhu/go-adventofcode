package int_set

type IntSet struct {
	data map[int]struct{}
}

func New() *IntSet {
	return &IntSet{data: make(map[int]struct{})}
}

func (s *IntSet) Add(b int) {
	s.data[b] = struct{}{}
}

func (s *IntSet) Remove(b int) {
	delete(s.data, b)
}

func (s *IntSet) IsMember(b int) bool {
	_, found := s.data[b]
	return found
}

func (s *IntSet) Intersection(o *IntSet) *IntSet {
	res := New()
	for k := range s.data {
		if o.IsMember(k) {
			res.Add(k)
		}
	}
	return res
}

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

func (s *IntSet) Each(fn func(int)) {
	for k := range s.data {
		fn(k)
	}
}
