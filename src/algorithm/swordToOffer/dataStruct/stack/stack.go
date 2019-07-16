package stack

type Stack struct {
	slice []int
}

func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
}

func (s *Stack) Pop() int{
	l := s.Size()
	if l == 0 {
		panic("stack empty")
	}

	val := s.slice[l-1]
	s.slice = s.slice[:l-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.slice)
}
