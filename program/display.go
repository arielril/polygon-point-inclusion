package program

import (
	"fmt"

	"github.com/arielril/basic-go-gl/util"
	"github.com/go-gl/gl/v2.1/gl"
)

var fps util.FPS

func displayFps() {
	acc := fps.SetFPS().GetAccumulated()
	if acc >= 1 { // print every second
		fmt.Printf("FPS: %v\n", fps.GetFPS())
		fps.Reset()
	}
}

func displayLine() {
	gl.PushMatrix()
	{
		gl.Color3f(1, 0, 0)
		gl.LineWidth(5)

		gl.Begin(gl.LINES)

		gl.Vertex2d(0, 0)
		gl.Vertex2d(10, 10)

		gl.End()
	}
	gl.PopMatrix()
}
