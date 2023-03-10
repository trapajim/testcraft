package testcraft

type Number interface {
	int | int32 | int64 | int8 | float32 | float64
}
type Sequence[T Number] struct {
	cur T
	inc T
}

func NewSequencer[T Number](start T) *Sequence[T] {
	return &Sequence[T]{
		cur: start,
		inc: 1,
	}
}

func (s *Sequence[T]) SetIncrement(inc T) {
	s.inc = inc
}

func (s *Sequence[T]) Next() T {
	v := s.cur
	s.cur += s.inc
	return v
}
