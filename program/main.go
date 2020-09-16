package program

import (
	"fmt"

	"github.com/arielril/basic-go-gl/algorithm"
	"github.com/arielril/basic-go-gl/object"
	"github.com/arielril/basic-go-gl/util"

	"github.com/go-gl/glfw/v3.3/glfw"
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
	parsedFile := util.ParseFile("./files/polygon1.txt")
	fileObject = object.NewObjectFromFile(parsedFile)
	randomPoints = object.GenerateRandomPoints(20)

	stripes = createStripes()
	setEdgeToStripes(fileObject.GetLines(), stripes)

	convexHull = object.NewPolygonFromPoints(
		algorithm.NewConvexHull(fileObject.GetPoints()),
	)
}

// Display the game
func Display(w *glfw.Window) {
	// displayLine()
	// displayFps()
	displayStripes()
	displayRandomPoints()
	displayFileObject()
	displayConvexHull()
}
