package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glu"
	// glmath "github.com/go-gl/mathgl/mgl64"
	"math"
)

const (
	Width  = 800
	Height = 600
	Title  = "Voxel Party"
)

func block() {
	// gl.ShadeModel(gl.FLAT)
	gl.Normal3d(0.0, 0.0, 1.0)

	gl.Begin(gl.LINE)
	gl.Color4f(1.0, 0.0, 0.0, 1.0)
	gl.Vertex3f(0.5, -0.5, -0.5)
	gl.Vertex3f(0.5, 0.5, -0.5)
	gl.End()

	gl.Begin(gl.QUADS)

	gl.Color4f(1.0, 1.0, 1.0, 1.0)
	gl.Vertex3f(0.5, -0.5, -0.5)
	gl.Vertex3f(0.5, 0.5, -0.5)
	gl.Vertex3f(-0.5, 0.5, -0.5)
	gl.Vertex3f(-0.5, -0.5, -0.5)

	gl.Vertex3f(0.5, -0.5, 0.5)
	gl.Vertex3f(0.5, 0.5, 0.5)
	gl.Vertex3f(-0.5, 0.5, 0.5)
	gl.Vertex3f(-0.5, -0.5, 0.5)

	gl.Vertex3f(0.5, -0.5, -0.5)
	gl.Vertex3f(0.5, 0.5, -0.5)
	gl.Vertex3f(0.5, 0.5, 0.5)
	gl.Vertex3f(0.5, -0.5, 0.5)

	gl.Vertex3f(-0.5, -0.5, 0.5)
	gl.Vertex3f(-0.5, 0.5, 0.5)
	gl.Vertex3f(-0.5, 0.5, -0.5)
	gl.Vertex3f(-0.5, -0.5, -0.5)

	gl.Vertex3f(0.5, 0.5, 0.5)
	gl.Vertex3f(0.5, 0.5, -0.5)
	gl.Vertex3f(-0.5, 0.5, -0.5)
	gl.Vertex3f(-0.5, 0.5, 0.5)

	gl.Vertex3f(0.5, -0.5, -0.5)
	gl.Vertex3f(0.5, -0.5, 0.5)
	gl.Vertex3f(-0.5, -0.5, 0.5)
	gl.Vertex3f(-0.5, -0.5, -0.5)

	gl.End()
}

var (
	positionX     = 1.0
	positionY     = 0.0
	positionZ     = -10.0
	positionSpeed = 1.0

	rotationX     = math.Pi / 2
	rotationY     = math.Pi / 2
	rotationSpeed = math.Pi / 32
)

var (
	block1 uint
)

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LoadIdentity()

	point := []float64{math.Cos(rotationY) * math.Sin(rotationX), math.Cos(rotationX), math.Sin(rotationY) * math.Sin(rotationX)}

	glu.LookAt(positionX, positionY, positionZ, positionX+point[0], positionY+point[1], positionZ+point[2], 0, 1, 0)

	gl.PushMatrix()
	gl.CallList(block1)
	gl.PopMatrix()
}

func key(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch glfw.Key(k) {
	case glfw.KeyW:
		positionZ += positionSpeed
	case glfw.KeyS:
		positionZ -= positionSpeed
	case glfw.KeyA:
		positionX += positionSpeed
	case glfw.KeyD:
		positionX -= positionSpeed
	case glfw.KeyQ:
		positionY += positionSpeed
	case glfw.KeyE:
		positionY -= positionSpeed
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	case glfw.KeyUp:
		rotationX -= rotationSpeed
	case glfw.KeyDown:
		rotationX += rotationSpeed
	case glfw.KeyLeft:
		rotationY -= rotationSpeed
	case glfw.KeyRight:
		rotationY += rotationSpeed
	default:
		return
	}
}

func reshape(window *glfw.Window, width, height int) {
	h := float64(height) / float64(width)

	znear := 5.0
	zfar := 30.0
	xmax := znear * 0.5

	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-xmax, xmax, -xmax*h, xmax*h, znear, zfar)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translated(0.0, 0.0, -20.0)
}

func Init() {
	pos := []float32{5.0, 5.0, 10.0, 0.0}
	red := []float32{0.8, 0.8, 0.8, 1.0}

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	gl.Lightfv(gl.LIGHT0, gl.POSITION, pos)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)

	block1 = gl.GenLists(1)
	gl.NewList(block1, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, red)
	block()
	gl.EndList()

	gl.Enable(gl.NORMALIZE)
}

func main() {
	if !glfw.Init() {
		panic("Failed to initialize GLFW")
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.DepthBits, 16)

	window, err := glfw.CreateWindow(Width, Height, Title, nil, nil)

	if err != nil {
		panic(err)
	}

	window.SetFramebufferSizeCallback(reshape)
	window.SetKeyCallback(key)

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	width, height := window.GetFramebufferSize()
	reshape(window, width, height)

	Init()

	for !window.ShouldClose() {
		draw()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
