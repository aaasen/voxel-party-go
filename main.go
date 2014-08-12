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

func grid(x float32, y float32, z float32, d float32, n float32) {
	width := d / n

	gl.Begin(gl.LINES)
	gl.Color3f(1.0, 1.0, 1.0)

	for i := -n; i <= n; i += 1.0 {
		gl.Vertex3f(x+width*i, y, z-d)
		gl.Vertex3f(x+width*i, y, z+d)

		gl.Vertex3f(x-d, y, z+width*i)
		gl.Vertex3f(x+d, y, z+width*i)
	}

	gl.End()
}

var (
	initialPosition = glmath.Vec3{0.0, 0.0, -10.0}
	camera          = NewCamera(initialPosition)
)

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LoadIdentity()

	position := camera.GetPosition()
	target := camera.GetTarget()

	glu.LookAt(position.X(), position.Y(), position.Z(), target.X(), target.Y(), target.Z(), 0, 1, 0)

	chunkManager.update(position)

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
	listManager.draw()
	gl.PopMatrix()
}

var (
	forward  = false
	backward = false
	left     = false
	right    = false
	up       = false
	down     = false

	listManager  = NewDisplayListManager()
	chunkManager = NewChunkManager(listManager)
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

	znear := 1.0
	zfar := 100.0
	xmax := znear * 1.0

	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-xmax, xmax, -xmax*h, xmax*h, znear, zfar)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func Init() {
	pos := []float32{0.0, 0.0, -10.0, 0.0}
	// red := []float32{0.8, 0.8, 0.8, 1.0}

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	gl.Lightfv(gl.LIGHT0, gl.POSITION, pos)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)

	n := 1

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			for z := 0; z < n; z++ {

				// position := glmath.Vec3{float64(x), float64(y), float64(z)}

				// listManager.add(NewChunk(position.Mul(16.0)))
			}
		}
	}

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
