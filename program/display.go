package program

import (
	"fmt"

	"github.com/arielril/polygon-point-inclusion/object"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var runningBenchmark bool = false

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
		for _, p := range randomPoints {
			if p.Raw().Inside {
				if p.Raw().Where == object.PointWhere.Polygon {
					gl.Color3f(0, 0, 1)
				} else if p.Raw().Where == object.PointWhere.ConvexHull {
					gl.Color3f(0.90, 0.90, 0.00)
				}
			} else {
				gl.Color3f(1, 0, 0)
			}
			gl.PointSize(13)

			gl.Begin(gl.POINTS)

			p.Draw()

			gl.End()
		}
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

func displayConvexHull() {
	gl.PushMatrix()
	{
		gl.Color3f(0.80, 0.27, 0.00)
		gl.LineWidth(4)

		objLines := convexHull.GetLines()
		gl.Begin(gl.LINES)

		for _, l := range objLines {
			l.Draw()
		}

		gl.End()
	}
	gl.PopMatrix()
}

// Display the game
func Display(w *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	if !runningBenchmark {
		// displayLine()
		// displayFps()
		displayStripes()
		displayRandomPoints()
		displayFileObject()
		displayConvexHull()
	} else {
		RunBenchmark()
		SetRunningBenchmark(false)
	}
}
