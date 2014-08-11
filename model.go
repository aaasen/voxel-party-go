package main

import (
	"github.com/go-gl/gl"
)

const (
	red   = iota
	blue  = iota
	green = iota
)

type Block struct {
	active    bool
	blockType int
}

func (*Block) draw() {
	gl.Begin(gl.QUADS)

	// when looking down the z axis:
	// front face
	gl.Normal3d(0.0, 0.0, -1.0)
	gl.Vertex3f(1.0, 0.0, 0.0)
	gl.Vertex3f(1.0, 1.0, 0.0)
	gl.Vertex3f(0.0, 1.0, 0.0)
	gl.Vertex3f(0.0, 0.0, 0.0)

	// back face
	gl.Normal3d(0.0, 0.0, -1.0)
	gl.Vertex3f(1.0, 0.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(0.0, 1.0, 1.0)
	gl.Vertex3f(0.0, 0.0, 1.0)

	// right face
	gl.Normal3d(1.0, 0.0, 1.0)
	gl.Vertex3f(1.0, 0.0, 0.0)
	gl.Vertex3f(1.0, 1.0, 0.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 0.0, 1.0)

	// left face
	gl.Normal3d(-1.0, 0.0, 1.0)
	gl.Vertex3f(0.0, 0.0, 1.0)
	gl.Vertex3f(0.0, 1.0, 1.0)
	gl.Vertex3f(0.0, 1.0, 0.0)
	gl.Vertex3f(0.0, 0.0, 0.0)

	// top face
	gl.Normal3d(0.0, 1.0, 0.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 0.0)
	gl.Vertex3f(0.0, 1.0, 0.0)
	gl.Vertex3f(0.0, 1.0, 1.0)

	// bottom face
	gl.Normal3d(0.0, -1.0, 0.0)
	gl.Vertex3f(1.0, 0.0, 0.0)
	gl.Vertex3f(1.0, 0.0, 1.0)
	gl.Vertex3f(0.0, 0.0, 1.0)
	gl.Vertex3f(0.0, 0.0, 0.0)

	gl.End()
}

func (block *Block) lighting() bool {
	return true
}

const (
	chunkWidth  = 16
	chunkHeight = 16
	chunkDepth  = 16
)

type Chunk struct {
	blocks [chunkWidth][chunkWidth][chunkWidth]Block
}

func NewChunk() *Chunk {
	return &Chunk{
		blocks: makeBlocks(),
	}
}

func (chunk *Chunk) draw() {
	gl.Color4f(1.0, 1.0, 1.0, 1.0)

	for x := 0; x < chunkWidth; x++ {
		for y := 0; y < chunkHeight; y++ {
			for z := 0; z < chunkDepth; z++ {
				gl.PushMatrix()

				gl.Translatef(float32(x), float32(y), float32(z))

				cube(float32(x), float32(y), float32(z))

				gl.PopMatrix()
			}
		}
	}
}

func (chunk *Chunk) lighting() bool {
	return true
}

func makeBlocks() [chunkWidth][chunkWidth][chunkWidth]Block {
	var blocks [chunkWidth][chunkHeight][chunkDepth]Block

	for x := 0; x < chunkWidth; x++ {
		for y := 0; y < chunkHeight; y++ {
			for z := 0; z < chunkDepth; z++ {

			}
		}
	}

	return blocks
}

// func forEach()
