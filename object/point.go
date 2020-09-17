package object

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
)

var rnd *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

type point struct {
	X, Y, Z float32
	Inside  bool
	Where   string // POLYGON | CONVEX_HULL
}

// Point interface
type Point interface {
	Raw() *point
	Draw()
	SetInside(where string) Point
}

// PointWhere is where is the point
var PointWhere = &struct {
	Polygon, ConvexHull string
}{
	Polygon:    "POLYGON",
	ConvexHull: "CONVEX_HULL",
}

func newPoint(x, y, z float32) Point {
	return &point{
		X:      x,
		Y:      y,
		Z:      z,
		Inside: false,
		Where:  "",
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

func (p *point) SetInside(where string) Point {
	if where == "" {
		p.Inside = false
		p.Where = ""
	} else {
		p.Inside = true
		p.Where = where
	}

	return p
}
