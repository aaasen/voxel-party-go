package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

func block() {
	gl.ShadeModel(gl.FLAT)
	gl.Normal3d(0.0, 0.0, 1.0)

	gl.Begin(gl.QUADS)

	gl.Vertex3f(-1, -1, 0)
	gl.Vertex3f(1, -1, 0)
	gl.Vertex3f(1, 1, 0)
	gl.Vertex3f(-1, 1, 0)

	gl.End()
}

var (
	rotationX float32 = 0.0
	rotationY float32 = 0.0
	rotationZ float32 = 0.0
	block1    uint
	angle     = 0.0
)

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Rotatef(rotationX, 1.0, 0.0, 0.0)
	gl.Rotatef(rotationY, 0.0, 1.0, 0.0)
	gl.Rotatef(rotationZ, 0.0, 0.0, 1.0)

	gl.PushMatrix()
	gl.Translated(0.0, 0.0, 0.0)
	gl.CallList(block1)
	gl.PopMatrix()

	gl.PopMatrix()
}

// change view angle, exit upon ESC
func key(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch glfw.Key(k) {
	case glfw.KeyZ:
		if mods&glfw.ModShift != 0 {
			rotationZ -= 5.0
		} else {
			rotationZ += 5.0
		}
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	case glfw.KeyUp:
		rotationX += 5.0
	case glfw.KeyDown:
		rotationX -= 5.0
	case glfw.KeyLeft:
		rotationY += 5.0
	case glfw.KeyRight:
		rotationY -= 5.0
	default:
		return
	}
}

// new window size
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

// program & OpenGL initialization
func Init() {
	pos := []float32{5.0, 5.0, 10.0, 0.0}
	red := []float32{0.8, 0.1, 0.0, 1.0}
	// green := []float32{0.0, 0.8, 0.2, 1.0}
	// blue := []float32{0.2, 0.2, 1.0, 1.0}

	gl.Lightfv(gl.LIGHT0, gl.POSITION, pos)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)

	// make the gears
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

	window, err := glfw.CreateWindow(800, 600, "Voxel Party", nil, nil)
	if err != nil {
		panic(err)
	}

	// Set callback functions
	window.SetFramebufferSizeCallback(reshape)
	window.SetKeyCallback(key)

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	width, height := window.GetFramebufferSize()
	reshape(window, width, height)

	// Parse command-line options
	Init()

	// Main loop
	for !window.ShouldClose() {
		// Draw gears
		draw()

		// Swap buffers
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
