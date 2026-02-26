package stack

type StackInt struct {
	data[]int
}

func (s *StackInt) Push(x int) {
	s.data = append(s.data, x)
}

func (s *StackInt) Len() int {
	return len(s.data)
}

func (s *StackInt) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *StackInt) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	idx := len(s.data) - 1
	value := s.data[idx]
	s.data[idx] = 0
	s.data = s.data[:idx]
	return value, true
}

func (s *StackInt) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	idx := len(s.data) -1
	value := s.data[idx]
	return value, true
}

