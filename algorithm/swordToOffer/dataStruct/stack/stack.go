package stack

type Stack struct {
	slice []int
	top   int
}

func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
	s.top = val
}

func (s *Stack) Pop() int {
	l := s.Size()
	if l == 0 {
		panic("stack empty")
	}

	val := s.slice[l-1]
	s.slice = s.slice[:l-1]
	if l >= 2 {
		s.top = s.slice[l-2]
	} else {
		s.top = 0
	}

	return val
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) Top() int {
	return s.top
}
