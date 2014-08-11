package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glu"
	glmath "github.com/go-gl/mathgl/mgl64"
)

const (
	Title = "Voxel Party"
)

var (
	width   = 800
	height  = 600
	centerX = float64(width) / 2.0
	centerY = float64(height) / 2.0
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
	initialPosition = glmath.Vec3{0.0, 0.0, -10.0}
	camera          = NewCamera(initialPosition)
)

var (
	block1 uint
)

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LoadIdentity()

	position := camera.GetPosition()

	target := camera.GetTarget()

	glu.LookAt(position.X(), position.Y(), position.Z(), target.X(), target.Y(), target.Z(), 0, 1, 0)

	if forward && !backward {
		camera.MoveForward(1.0)
	} else if backward && !forward {
		camera.MoveForward(-1.0)
	}

	if right && !left {
		camera.MoveRight(-1.0)
	} else if left && !right {
		camera.MoveRight(1.0)
	}

	if up && !down {
		camera.MoveUp(1.0)
	} else if down && !up {
		camera.MoveUp(-1.0)
	}

	camera.Tick()

	gl.PushMatrix()
	gl.CallList(block1)
	gl.PopMatrix()
}

var (
	forward  = false
	backward = false
	left     = false
	right    = false
	up       = false
	down     = false
)

func mouse(window *glfw.Window, xpos, ypos float64) {
	window.SetCursorPosition(float64(width)/2.0, float64(height)/2.0)

	dx := xpos - centerX
	dy := ypos - centerY

	camera.Rotate(glmath.Vec2{dy, dx})
}

func key(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		switch glfw.Key(k) {
		case glfw.KeyW:
			forward = true
		case glfw.KeyS:
			backward = true
		case glfw.KeyA:
			left = true
		case glfw.KeyD:
			right = true
		case glfw.KeySpace:
			if mods == glfw.ModShift {
				down = true
			} else {
				up = true
			}
		case glfw.KeyEscape:
			window.SetShouldClose(true)
		}
	} else if action == glfw.Release {
		switch glfw.Key(k) {
		case glfw.KeyW:
			forward = false
		case glfw.KeyS:
			backward = false
		case glfw.KeyA:
			left = false
		case glfw.KeyD:
			right = false
		case glfw.KeySpace:
			if mods == glfw.ModShift {
				down = false
			} else {
				up = false
			}
		}
	}
}

func reshape(window *glfw.Window, width, height int) {
	width = width
	height = height
	centerX = float64(width) / 2.0
	centerY = float64(height) / 2.0

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
	pos := []float32{0.0, 0.0, -10.0, 0.0}
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

	window, err := glfw.CreateWindow(width, height, Title, nil, nil)

	if err != nil {
		panic(err)
	}

	window.SetFramebufferSizeCallback(reshape)
	window.SetCursorPositionCallback(mouse)
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
