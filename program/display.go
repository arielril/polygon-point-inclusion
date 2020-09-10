package program

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

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

func displayFileObject() {
	gl.PushMatrix()
	{
		gl.Color3f(1, 0, 1)
		gl.LineWidth(5)

		objLines := fileObject.GetLines()

		gl.Begin(gl.LINES)
		for _, l := range objLines {
			l.Draw()
		}
		gl.End()

	}
	gl.PopMatrix()
}

func displayRandomPoints() {
	gl.PushMatrix()
	{
		gl.Color3f(1, 0, 0)
		gl.PointSize(20)

		gl.Begin(gl.POINTS)

		for _, p := range randomPoints {
			p.Draw()
		}

		gl.End()
	}
	gl.PopMatrix()
}

func displayStripes() {
	if stripes == nil {
		return
	}

	gl.PushMatrix()
	{
		gl.Color3f(0, 1, 1)
		gl.LineWidth(3)

		gl.Begin(gl.LINES)

		for _, s := range stripes {
			s.Draw()
		}

		gl.End()
	}
	gl.PopMatrix()
}
