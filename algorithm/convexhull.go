package algorithm

import (
	"fmt"

	"github.com/arielril/basic-go-gl/object"
)

//https://www.geeksforgeeks.org/convex-hull-set-1-jarviss-algorithm-or-wrapping/

// getLeftMostPoint
func getLeftMostPoint(points []object.Point) int {
	minIndex := 0

	for i := 0; i < len(points); i++ {
		actualPoint := points[i].Raw()
		minimalPoint := points[minIndex].Raw()
		if actualPoint.X < minimalPoint.X {
			minIndex = i
		} else if actualPoint.X == minimalPoint.X {
			if actualPoint.Y > minimalPoint.Y {
				minIndex = i
			}
		}
	}

	return minIndex
}

// orientation
func orientation(p, q, r object.Point) string {
	pr := p.Raw()
	qr := q.Raw()
	rr := r.Raw()

	value := (qr.Y-pr.Y)*(rr.X-qr.X) -
		(qr.X-pr.X)*(rr.Y-qr.Y)

	if value == 0 {
		return "COLINEAR"
	} else if value > 0 {
		return "CLOCKWISE"
	} else {
		return "COUNTERCLOCKWISE"
	}
}

// NewConvexHull create a Convex Hull object
func NewConvexHull(points []object.Point) []object.Point {
	if len(points) < 3 {
		fmt.Println("not enough points to go")
		return nil
	}

	leftMostPointIndex := getLeftMostPoint(points)
	numberOfPoints := len(points)

	hull := make([]int, 0)

	actualPointIndex := leftMostPointIndex
	otherPointIndex := 0

	for {
		hull = append(hull, actualPointIndex)

		otherPointIndex = (actualPointIndex + 1) % numberOfPoints

		for i := 0; i < numberOfPoints; i++ {
			if orientation(
				points[actualPointIndex],
				points[i],
				points[otherPointIndex],
			) == "COUNTERCLOCKWISE" {
				otherPointIndex = i
			}
		}

		actualPointIndex = otherPointIndex

		if actualPointIndex == leftMostPointIndex {
			break
		}
	}

	hullPoints := make([]object.Point, 0)

	for _, idx := range hull {
		hullPoints = append(hullPoints, points[idx])
	}

	return hullPoints
}
