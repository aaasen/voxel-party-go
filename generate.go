package main

// Terrain generation

import ()

func generateChunk(position ChunkCoordinate) *Chunk {
	var blocks [chunkWidth][chunkHeight][chunkDepth]Block

	for x := 0; x < chunkWidth; x++ {
		for z := 0; z < chunkDepth; z++ {
			blocks[x][0][z] = Block{true, red}
		}
	}

	return &Chunk{
		blocks:   blocks,
		position: position,
	}
}
