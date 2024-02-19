package collections

type Set map[string]struct{}

// add an element to the set
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// remove an element from the set
func (s Set) Remove(elem string) {
	delete(s, elem)
}

// check if an element is in the set
func (s Set) Contains(elem string) bool {
	_, ok := s[elem]
	return ok
}

// get the size of the set
func (s Set) Size() int {
	return len(s)
}

// returns all keys
func (s Set) Keys() []string {
	keys := make([]string, 0, s.Size())

	for k := range s {
		keys = append(keys, k)
	}

	return keys
}
