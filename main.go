package main

import (
	"fmt"
	"runtime"

	"github.com/arielril/basic-go-gl/opengl"
	"github.com/arielril/basic-go-gl/program"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

const (
	wWidth  = 800
	wHeight = 600
)

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("failed to init glfw: %v", err))
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
}

func main() {
	win, _ := opengl.NewWindow(wWidth, wHeight, "Polygon Point Inclusion")
	defer glfw.Terminate()

	win.SetKeyCallback(opengl.KeyCallback)
	win.SetCharCallback(opengl.CharCallback)

	opengl.Setup()
	program.Init()

	/*
		1. Create base polygon
		2. Run brute force algorithm
		3. Create Convex Hull polygon
		4. Run the other algorithm
	*/

	for !win.ShouldClose() {
		opengl.Reshape(win)
		program.Display(win)

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
