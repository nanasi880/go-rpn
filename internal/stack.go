package internal

import "github.com/shopspring/decimal"

var (
	zero decimal.Decimal
)

type Stack struct {
	v []decimal.Decimal
}

func (s *Stack) Push(d decimal.Decimal) {
	s.v = append(s.v, d)
}

func (s *Stack) Pop() (decimal.Decimal, bool) {

	if len(s.v) == 0 {
		return zero, false
	}

	v := s.v[len(s.v)-1]

	s.v = s.v[:len(s.v)-1]

	return v, true
}

func (s *Stack) Pop2() (decimal.Decimal, decimal.Decimal, bool) {

	if len(s.v) < 2 {
		return zero, zero, false
	}

	v0, v1 := s.v[len(s.v)-1], s.v[len(s.v)-2]

	s.v = s.v[:len(s.v)-2]

	return v0, v1, true
}
