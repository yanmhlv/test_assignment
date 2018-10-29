package app

type (
	Stack struct {
		ary   []Pair
		count int
	}
)

func NewStack() *Stack { return &Stack{} }

func (s *Stack) Push(p Pair) {
	s.ary = append(s.ary[:s.count], p)
	s.count++
}

func (s *Stack) Pop() *Pair {
	if s.count == 0 {
		return nil
	}
	s.count--
	p := s.ary[s.count]
	return &p
}
