package object

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

type point struct {
	X, Y, Z float32
}

// Point interface
type Point interface {
	Raw() *point
	Draw()
}

func newPoint(x, y, z float32) Point {
	return &point{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewPoint2D creates a 2D point
func NewPoint2D(x, y float32) Point {
	return newPoint(x, y, 0)
}

func (p *point) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p.X, p.Y, p.Z)
}

// Raw returns the raw values from the point
func (p *point) Raw() *point {
	return p
}

// Draw draws the point in the OpenGL window
func (p *point) Draw() {
	gl.PushMatrix()
	{
		pr := p.Raw()
		gl.Vertex2f(pr.X, pr.Y)
	}
	gl.PopMatrix()
}
