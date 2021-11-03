package string_set

type StringSet struct {
	data map[string]struct{}
}

func New() *StringSet {
	return &StringSet{data: make(map[string]struct{})}
}

func (s *StringSet) Add(b string) {
	s.data[b] = struct{}{}
}

func (s *StringSet) Remove(b string) {
	delete(s.data, b)
}

func (s *StringSet) IsMember(b string) bool {
	_, found := s.data[b]
	return found
}

func (s *StringSet) Intersection(o *StringSet) *StringSet {
	res := New()
	for k := range s.data {
		if o.IsMember(k) {
			res.Add(k)
		}
	}
	return res
}

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

func (s *StringSet) Each(fn func(string)) {
	for k := range s.data {
		fn(k)
	}
}
