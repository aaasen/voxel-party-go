package main

import (
	"github.com/go-gl/gl"
	"math/rand"
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

func (block *Block) draw() {
	if block.active {
		gl.Begin(gl.QUADS)

		switch block.blockType {
		case red:
			gl.Color3f(1.0, 0.0, 0.0)
		case green:
			gl.Color3f(0.0, 1.0, 0.0)
		case blue:
			gl.Color3f(0.0, 0.0, 1.0)
		}

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
}

func (block *Block) lighting() bool {
	return false
}

const (
	chunkWidth  = 16
	chunkHeight = 16
	chunkDepth  = 16
)

type Chunk struct {
	blocks   [chunkWidth][chunkWidth][chunkWidth]Block
	position []float32
}

func NewChunk(position []float32) *Chunk {
	return &Chunk{
		blocks:   makeBlocks(),
		position: position,
	}
}

func (chunk *Chunk) draw() {
	gl.Color4f(1.0, 1.0, 1.0, 1.0)

	for x := 0; x < chunkWidth; x++ {
		for y := 0; y < chunkHeight; y++ {
			for z := 0; z < chunkDepth; z++ {
				gl.PushMatrix()
				gl.Translatef(chunk.position[0], chunk.position[1], chunk.position[2])

				gl.PushMatrix()
				gl.Translatef(float32(x), float32(y), float32(z))

				block := chunk.blocks[x][y][z]
				block.draw()

				gl.PopMatrix()

				gl.PopMatrix()
			}
		}
	}
}

func (chunk *Chunk) lighting() bool {
	return false
}

func makeBlocks() [chunkWidth][chunkWidth][chunkWidth]Block {
	var blocks [chunkWidth][chunkHeight][chunkDepth]Block

	for x := 0; x < chunkWidth; x++ {
		for y := 0; y < chunkHeight; y++ {
			for z := 0; z < chunkDepth; z++ {
				blockType := blue

				if x%2 == 0 {
					blockType = red
				}

				if z%2 == 0 {
					blockType = green
				}

				active := false

				if rand.Int31n(2) == 0 {
					active = true
				}

				blocks[x][y][z] = Block{active, blockType}
			}
		}
	}

	return blocks
}
