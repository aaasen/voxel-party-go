package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glmath "github.com/go-gl/mathgl/mgl64"
	"math"
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
	position ChunkCoordinate
}

func (chunk *Chunk) draw() {
	gl.Color4f(1.0, 1.0, 1.0, 1.0)

	for x := 0; x < chunkWidth; x++ {
		for y := 0; y < chunkHeight; y++ {
			for z := 0; z < chunkDepth; z++ {
				gl.PushMatrix()
				position := chunk.position.toVector()

				gl.Translated(position.X(), position.Y(), position.Z())

				gl.PushMatrix()
				gl.Translated(float64(x), float64(y), float64(z))

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

const (
	DefaultRenderDistance = 1
)

type ChunkManager struct {
	listManager    *DisplayListManager
	chunks         map[string]*Chunk
	ids			   map[string]uint
	renderDistance int
	lastPosition   ChunkCoordinate
}

type ChunkCoordinate []int

func (coord *ChunkCoordinate) Id() string {
	return fmt.Sprintf("%v:%v:%v", coord.X(), coord.Y(), coord.Z())
}

func (coord ChunkCoordinate) X() int {
	return coord[0]
}

func (coord ChunkCoordinate) Y() int {
	return coord[1]
}

func (coord ChunkCoordinate) Z() int {
	return coord[2]
}

func idsFromCoords(coords []ChunkCoordinate) []string {
	ids := make([]string, len(coords), len(coords))

	for i, coord := range coords {
		ids[i] = coord.Id()
	}

	return ids
}

func (coord ChunkCoordinate) Equals(other ChunkCoordinate) bool {
	return coord.X() == other.X() && coord.Y() == other.Y() && coord.Z() == other.Z()
}

func NewChunkManager(listManager *DisplayListManager) *ChunkManager {
	return &ChunkManager{
		listManager:    listManager,
		chunks:         make(map[string]*Chunk),
		ids: 			make(map[string]uint),
		renderDistance: DefaultRenderDistance,
		lastPosition:   nil,
	}
}

func (manager *ChunkManager) update(position glmath.Vec3) {
	chunkCoordinate := toChunkCoordinates(position)

	if manager.lastPosition == nil || !chunkCoordinate.Equals(manager.lastPosition) {
		chunksToLoad := chunksWithinDistance(chunkCoordinate, manager.renderDistance)

		for _, chunkCoord := range chunksToLoad {
			id := chunkCoord.Id()

			_, ok := manager.chunks[id]

			if !ok {
				chunk := generateChunk(chunkCoord)
				manager.chunks[id] = chunk

				listId := manager.listManager.add(chunk)
				manager.ids[id] = listId
			}
		}

		if manager.lastPosition != nil {
			chunksToDestroy := chunksWithinDistance(manager.lastPosition, manager.renderDistance)
			chunkIdsToLoad := idsFromCoords(chunksToLoad)

			for _, toDestroy := range chunksToDestroy {
				id := toDestroy.Id()

				if !stringInSlice(id, chunkIdsToLoad) {
					delete(manager.chunks, id)

					listId := manager.ids[id]
					manager.listManager.remove(listId)
					delete(manager.ids, id)
				}
			}
		}

		manager.lastPosition = chunkCoordinate
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func chunksWithinDistance(position ChunkCoordinate, distance int) []ChunkCoordinate {
	chunks := make([]ChunkCoordinate, 0, distance*distance*distance)

	for x := -distance; x <= distance; x++ {
		for z := -distance; z <= distance; z++ {
			chunks = append(chunks, ChunkCoordinate{position.X() + x, position.Y(), position.Z() + z})
		}
	}

	return chunks
}

func toChunkCoordinates(position glmath.Vec3) ChunkCoordinate {
	return ChunkCoordinate{floorInt(position.X() / chunkWidth),
		floorInt(position.Y() / chunkHeight),
		floorInt(position.Z() / chunkDepth)}
}

func (position *ChunkCoordinate) toVector() glmath.Vec3 {
	return glmath.Vec3{float64(position.X() * chunkWidth), float64(position.Y() * chunkHeight), float64(position.Z() * chunkDepth)}
}

func floorInt(x float64) int {
	return int(math.Floor(x))
}
