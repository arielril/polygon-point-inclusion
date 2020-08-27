package program

import (
	"fmt"

	"github.com/arielril/basic-go-gl/object"
	"github.com/arielril/basic-go-gl/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var fps util.FPS
var fileObject object.Polygon
var randomPoints []object.Point

// Init the game
func Init() {
	fmt.Println("Init game")
	fps = util.NewFps()
	parsedFile := util.ParseFile("./files/polygon1.txt")
	fileObject = object.NewObjectFromFile(parsedFile)
	randomPoints = object.GenerateRandomPoints(20)
}

// Display the game
func Display(w *glfw.Window) {
	// displayLine()
	// displayFps()
	displayRandomPoints()
	displayFileObject()
}
