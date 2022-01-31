package methodset

type (
	MethodSet map[string]string
)

func (s *MethodSet) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *MethodSet) Remove(key string) {
	delete(*s, key)
}

func (s *MethodSet) Clear() {
	for k := range *s {
		delete(*s, k)
	}
}
func (s *MethodSet) Copy() MethodSet {
	newMap := make(MethodSet)
	for k, v := range *s {
		newMap[k] = v
	}
	return newMap
}
func (s *MethodSet) Add(key string, value string) {
	(*s)[key] = value
}
