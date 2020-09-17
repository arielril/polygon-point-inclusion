package algorithm

import (
	"github.com/arielril/polygon-point-inclusion/object"
)

// RunBruteForce runs the brute force algorithm; returns the quantity of call to Intersect2D
func RunBruteForce(pointList []object.Point, polygonLines []object.Line) int {
	intersectCalls := 0
	for _, p := range pointList {
		lineP1 := object.NewPoint2D(p.Raw().X, p.Raw().Y)
		lineP2 := object.NewPoint2D(10, p.Raw().Y)
		stepLine := object.NewLine(lineP1, lineP2)

		stepLineIntersections := 0

		// inside if
		// 1. count of intersections is odd
		// 2. point lies on an edge of polygon
		for _, pl := range polygonLines {
			intersectCalls++

			if stepLine.Intersect2D(pl) {
				stepLineIntersections++
			}
		}

		if (stepLineIntersections % 2) != 0 {
			p.SetInside(object.PointWhere.Polygon)
		}
	}

	return intersectCalls
}
