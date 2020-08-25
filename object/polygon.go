package object

import "github.com/arielril/basic-go-gl/util"

// Polygon interface
type Polygon interface {
	GetPoints() []Point
	GetLines() []Line
}

type polygon struct {
	points []Point
	lines  []Line
}

func createPolygonPoints(poly *polygon, parsedFile util.ParsedFile) {
	poly.points = make([]Point, 0)

	parsedFile = parsedFile[1:]
	for _, p := range parsedFile {
		point := NewPoint2D(p[0], p[1])
		poly.points = append(poly.points, point)
	}
}

func createPolygonLines(poly *polygon) {
	for i := 0; i < len(poly.points); i++ {
		p1 := poly.points[i]
		p2 := poly.points[(i+1)%len(poly.points)]
		l := NewLine(p1, p2)

		poly.lines = append(poly.lines, l)
	}
}

// NewObjectFromFile create an object from a parsed file
func NewObjectFromFile(parsedFile util.ParsedFile) Polygon {
	poly := &polygon{}

	createPolygonPoints(poly, parsedFile)
	createPolygonLines(poly)

	return poly
}

func (p *polygon) GetPoints() []Point {
	return p.points
}

func (p *polygon) GetLines() []Line {
	return p.lines
}
