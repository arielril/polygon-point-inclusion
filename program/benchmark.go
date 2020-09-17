package program

import (
	"fmt"
	"time"

	"github.com/arielril/polygon-point-inclusion/algorithm"
	"github.com/arielril/polygon-point-inclusion/object"
)

var benchPointQuantity = []int{200, 2000, 20000}

func _resetPointInsideWhere(points []object.Point) {
	for _, point := range points {
		point.SetInside("")
	}
}

type run struct {
	Points  int
	Elapsed time.Duration
	Calls   int
}

func (r *run) String() string {
	return fmt.Sprintf("Points\t%v\tElapsed\t%v\tCalls\t%v", r.Points, r.Elapsed, r.Calls)
}

type benchResult map[string][]*run

// RunBenchmark do
func RunBenchmark() {
	result := benchResult{
		"BRUTE_FORCE": make([]*run, 0),
		"CONVEX_HULL": make([]*run, 0),
	}

	for _, pointQty := range benchPointQuantity {

		randomPoints = object.GenerateRandomPoints(pointQty)

		startTime := time.Now()
		calls := algorithm.RunBruteForce(randomPoints, fileObject.GetLines())
		elapsed := time.Now().Sub(startTime)

		result["BRUTE_FORCE"] = append(result["BRUTE_FORCE"], &run{
			Points:  pointQty,
			Elapsed: elapsed,
			Calls:   calls,
		})

		fmt.Printf("Bruteforce (%v) - %v\n", pointQty, elapsed)

		_resetPointInsideWhere(randomPoints)

		startTime = time.Now()
		calls = algorithm.RunConvexHull(randomPoints, convexHull)
		elapsed = time.Now().Sub(startTime)

		result["CONVEX_HULL"] = append(result["CONVEX_HULL"], &run{
			Points:  pointQty,
			Elapsed: elapsed,
			Calls:   calls,
		})

		fmt.Printf("Convex Hull (%v) - %v\n", pointQty, elapsed)
	}
	fmt.Println()

	_showResult(result)

	randomPoints = object.GenerateRandomPoints(200)
}

func _showResult(r benchResult) {
	fmt.Println()
	fmt.Println("Legend: BF = Brute Force | CH = Convex Hull")
	fmt.Println()
	fmt.Println("Points\t;BF Time\t;BF Call Intersection\t;CH Time\t;CH Calls Vector Prod\t;")

	for i, pointQtd := range benchPointQuantity {
		bruteRes := r["BRUTE_FORCE"][i]
		hullRes := r["CONVEX_HULL"][i]

		fmt.Printf(
			"%v\t;%v\t;%v\t\t\t;%v\t;%v\t\t\t;\n",
			pointQtd,
			bruteRes.Elapsed,
			bruteRes.Calls,
			hullRes.Elapsed,
			hullRes.Calls,
		)
	}

	fmt.Println()
}
