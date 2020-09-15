package program

import (
	"fmt"

	"github.com/arielril/basic-go-gl/algorithm"
	"github.com/arielril/basic-go-gl/object"
	"github.com/arielril/basic-go-gl/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var fps util.FPS
var fileObject object.Polygon
var randomPoints []object.Point
var stripes []object.Line
var convexHull object.Polygon

// Init the game
func Init() {
	fmt.Println("Init game")
	fps = util.NewFps()
	parsedFile := util.ParseFile("./files/polygon1.txt")
	fileObject = object.NewObjectFromFile(parsedFile)
	randomPoints = object.GenerateRandomPoints(20)

	stripes = createStripes()
	setPointsToStripes(randomPoints, stripes)

	convexHull = object.NewPolygonFromPoints(
		algorithm.NewConvexHull(fileObject.GetPoints()),
	)
}

func setPointsToStripes(rndPoints []object.Point, stripeList []object.Line) {
	for _, p := range rndPoints {
		for _, s := range stripeList {
			fmt.Printf("point > %v\nstripe > %v\n", p, s)
		}
	}
}

func createStripes() []object.Line {
	const minY = 0
	const maxY = 10

	nStripes := make([]object.Line, maxY+1)

	for i := minY; i <= maxY; i++ {
		p1 := object.NewPoint2D(0, float32(i))
		p2 := object.NewPoint2D(10, float32(i))
		line := object.NewLine(p1, p2)

		nStripes[i] = line
	}

	return nStripes
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
