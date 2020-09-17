package program

import (
	"fmt"

	"github.com/arielril/polygon-point-inclusion/algorithm"
	"github.com/arielril/polygon-point-inclusion/object"
	"github.com/arielril/polygon-point-inclusion/util"
)

const stripeMinY = 0
const stripeMaxY = 10

var fps util.FPS
var fileObject object.Polygon
var randomPoints []object.Point
var stripes []object.Stripe
var convexHull object.Polygon

// Init the game
func Init() {
	fmt.Println("Init program...")
	fps = util.NewFps()
	parsedFile := util.ParseFile("./files/polygon2.txt")
	fileObject = object.NewObjectFromFile(parsedFile)
	randomPoints = object.GenerateRandomPoints(2000)

	stripes = createStripes()
	setEdgeToStripes(fileObject.GetLines(), stripes)

	convexHull = object.NewPolygonFromPoints(
		algorithm.NewConvexHull(fileObject.GetPoints()),
	)

	// show the initial colors of the points
	algorithm.RunConvexHull(randomPoints, convexHull)
	algorithm.RunBruteForce(randomPoints, fileObject.GetLines())
}
