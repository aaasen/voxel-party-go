package main

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
)

func glfwErrorCallback(err glfw.ErrorCode, description string) {
	fmt.Println("%v: %v", err, description)
}

const (
	width  = 640
	height = 400
	title  = "Voxel Party"
)

func main() {
	fmt.Println("Hello, world")

	glfw.SetErrorCallback(glfwErrorCallback)

	if !glfw.Init() {
		panic("cannot initialize GLFW")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		//Do OpenGL stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
