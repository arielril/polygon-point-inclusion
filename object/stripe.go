package object

import "github.com/go-gl/gl/v2.1/gl"

type stripe struct {
	Bottom, Top float32
	Lines       []Line
}

// Stripe interface
type Stripe interface {
	Raw() *stripe
	Draw()
	AddLine(l Line)
}

// NewStripe create a stripe
func NewStripe(bottom, top float32) Stripe {
	s := &stripe{
		Bottom: bottom,
		Top:    top,
		Lines:  make([]Line, 0),
	}

	return s
}

func (s *stripe) Raw() *stripe {
	return s
}

func (s *stripe) AddLine(l Line) {
	s.Lines = append(s.Lines, l)
}

func (s *stripe) Draw() {
	gl.PushMatrix()
	{
		gl.Vertex2f(0, s.Bottom)
		gl.Vertex2f(10, s.Bottom)

		gl.Vertex2f(0, s.Top)
		gl.Vertex2f(10, s.Top)
	}
	gl.PopMatrix()
}
