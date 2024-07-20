package sandbox

import "errors"

type Stack struct {
	data []string
}

func (s *Stack) Push(data string) {
	s.data = append(s.data, data)
}

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		topIdx := len(s.data) - 1
		s.data = s.data[:topIdx]
	}
}

func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}

	return s.data[len(s.data)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
