package object

import "github.com/go-gl/gl/v2.1/gl"

type line struct {
	P1, P2 Point
}

// Line interface
type Line interface {
	Intersect2D(l Line) bool
	Raw() *line
	Draw()
}

// NewLine creates a new line from two points
func NewLine(p1, p2 Point) Line {
	return &line{
		P1: p1,
		P2: p2,
	}
}

// Raw returns the raw line
func (l *line) Raw() *line {
	return l
}

// _intersec2D Compute the intersection between 2 lines
/* ********************************************************************** */
/*                                                                        */
/*  Calcula a interseccao entre 2 retas (no plano "XY" Z = 0)             */
/*                                                                        */
/* k : ponto inicial da reta 1                                            */
/* l : ponto final da reta 1                                              */
/* m : ponto inicial da reta 2                                            */
/* n : ponto final da reta 2                                              */
/*                                                                        */
/* s: valor do parâmetro no ponto de interseção (sobre a reta KL)         */
/* t: valor do parâmetro no ponto de interseção (sobre a reta MN)         */
/*                                                                        */
/* ********************************************************************** */
func _intersect2D(k, l, m, n Point) (s, t float32, ok bool) {
	pk := k.Raw()
	pl := l.Raw()
	pm := m.Raw()
	pn := n.Raw()

	det := (pn.X-pm.X)*(pl.Y-pk.Y) - (pn.Y-pm.Y)*(pl.X-pk.X)

	if det == 0 {
		return 0, 0, false // no intersection
	}

	s = ((pn.X-pm.X)*(pm.Y-pk.Y) - (pn.Y-pm.Y)*(pm.X-pk.X)) / det
	t = ((pl.X-pk.X)*(pm.Y-pk.Y) - (pl.Y-pk.Y)*(pm.X-pk.X)) / det

	return s, t, true
}

// Intersect2D returns if two lines have an intersection
func (l *line) Intersect2D(l1 Line) bool {
	l1Raw := l1.Raw()

	s, t, ok := _intersect2D(
		l.P1, l.P2,
		l1Raw.P1, l1Raw.P2,
	)

	if ok && (s >= 0 && s <= 1 && t >= 0 && t <= 1) {
		return true
	}

	return false
}

func (l *line) Draw() {
	gl.PushMatrix()
	{
		p1r := l.P1.Raw()
		p2r := l.P2.Raw()

		gl.Vertex2f(p1r.X, p1r.Y)
		gl.Vertex2f(p2r.X, p2r.Y)
	}
	gl.PopMatrix()
}
