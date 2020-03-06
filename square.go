package shape

type Square struct {
  side float64
}

func NewSquare(sidex float64) *Square {
  return &Square{
    side: sidex,
  }
}

func (s *Square) Perimeter() float64 {
  return 4 * s.side
}
