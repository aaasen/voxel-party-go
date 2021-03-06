package main

import (
	glmath "github.com/go-gl/mathgl/mgl64"
	"math"
)

var (
	forwardUnit = glmath.Vec3{1.0, 0.0, 0.0}
	rightUnit   = glmath.Vec3{0.0, 0.0, 1.0}
	upUnit      = glmath.Vec3{0.0, 1.0, 0.0}
)

type Camera struct {
	position     glmath.Vec3
	velocity     glmath.Vec3
	acceleration glmath.Vec3

	speed float64
	drag  float64

	targetUnit glmath.Vec3
	target     glmath.Vec3
	rotation   glmath.Vec2

	rotationSpeed float64
}

func NewCamera(position glmath.Vec3) *Camera {
	camera := &Camera{
		position:     position,
		velocity:     glmath.Vec3{0.0, 0.0, 0.0},
		acceleration: glmath.Vec3{0.0, 0.0, 0.0},

		speed: 0.5,
		drag:  0.5,

		targetUnit: glmath.Vec3{0.0, 0.0, 0.0},
		target:     glmath.Vec3{0.0, 0.0, 0.0},
		rotation:   glmath.Vec2{math.Pi / 2.0, math.Pi / 2.0},

		rotationSpeed: 0.001,
	}

	camera.Tick()

	return camera
}

func (camera *Camera) MoveForward(amount float64) {
	camera.acceleration = camera.acceleration.Add(camera.targetUnit.Mul(amount * camera.speed))
}

func (camera *Camera) MoveRight(amount float64) {
	camera.acceleration = camera.acceleration.Sub(camera.targetUnit.Cross(upUnit).Mul(amount * camera.speed))
}

func (camera *Camera) MoveUp(amount float64) {
	camera.acceleration = camera.acceleration.Add(upUnit.Mul(amount * camera.speed))
}

func (camera *Camera) Rotate(vec glmath.Vec2) {
	camera.rotation = camera.rotation.Add(vec.Mul(camera.rotationSpeed))
}

// Tick updates all fields of the Camera
func (camera *Camera) Tick() {
	camera.velocity = camera.velocity.Add(camera.acceleration).Mul(camera.drag)

	camera.acceleration = camera.acceleration.Mul(0.0)

	camera.position = camera.position.Add(camera.velocity)

	camera.targetUnit = camera.calculateTargetUnit()

	camera.target = camera.position.Add(camera.targetUnit)
}

func (camera *Camera) GetPosition() glmath.Vec3 {
	return camera.position
}

func (camera *Camera) calculateTargetUnit() glmath.Vec3 {
	x := camera.rotation.X()
	y := camera.rotation.Y()

	return glmath.Vec3{math.Cos(y) * math.Sin(x), math.Cos(x), math.Sin(y) * math.Sin(x)}
}

func (camera *Camera) GetTarget() glmath.Vec3 {
	return camera.target
}
