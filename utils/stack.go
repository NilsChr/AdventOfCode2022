package utils

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(data T) {
	*s = append(*s, data)
}

func (s *Stack[T]) Pop() (T, bool) {
	var result T
	if s.IsEmpty() {
		return result, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack[T]) Peak() (T, bool) {
	var result T
	if s.IsEmpty() {
		return result, false
	} else {
		return (*s)[len(*s)-1], true
	}
}
