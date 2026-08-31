package stack

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		return *new(T), false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		return *new(T), false
	}

	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) IsEmpty() bool { return len(s.elements) == 0 }

func (s *Stack[T]) Size() int { return len(s.elements) }
