package generics

//
//type StackOfInts struct {
//	values []int
//}
//
//func (s *StackOfInts) Push(value int) {
//	s.values = append(s.values, value)
//}
//
//func (s *StackOfInts) IsEmpty() bool {
//	return len(s.values) == 0
//}
//
//func (s *StackOfInts) Pop() (int, bool) {
//	if s.IsEmpty() {
//		return 0, false
//	}
//
//	index := len(s.values) - 1
//	el := s.values[index]
//	s.values = s.values[:index]
//	return el, true
//}
//
//type StackOfStrings struct {
//	values []string
//}
//
//func (s *StackOfStrings) Push(value string) {
//	s.values = append(s.values, value)
//}
//
//func (s *StackOfStrings) IsEmpty() bool {
//	return len(s.values) == 0
//}
//
//func (s *StackOfStrings) Pop() (string, bool) {
//	if s.IsEmpty() {
//		return "", false
//	}
//
//	index := len(s.values) - 1
//	el := s.values[index]
//	s.values = s.values[:index]
//	return el, true
//}

type StackOfInts = Stack
type StackOfStrings = Stack

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var zero interface{}
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
