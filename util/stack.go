package util

type Stack[T any] struct {
	data   []T
	length int
}

func NewStack[T any]() Stack[T] {
	data := make([]T, 4)
	return Stack[T]{length: 0, data: data}
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) Push(element T) {
	if s.length >= len(s.data) {
		newData := make([]T, 2*(len(s.data))+1)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.length] = element
	s.length++
}

func (s *Stack[T]) Peek() T {
	if s.length == 0 {
		panic("stack is empty")
	}

	return s.data[s.length-1]
}

func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		panic("stack is empty")
	}

	s.length--
	return s.data[s.length]
}

func (s *Stack[T]) Reverse() {
	ReverseSlice(s.data[0:s.length])
}
