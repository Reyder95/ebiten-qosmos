package ebitenqosmosclasses

type Chunk struct {
	TileList []Tile
	Loaded   bool
	ChunkId  int
}

func (c *Chunk) SetChunk(tileList []Tile, chunkId int) {
	c.TileList = tileList
	c.ChunkId = chunkId
}

func (c *Chunk) LoadChunk() {
	c.Loaded = true
}

func (c *Chunk) UnloadChunk() {
	c.Loaded = false
}
