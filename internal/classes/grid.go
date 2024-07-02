package ebitenqosmosclasses

type Grid struct {
	CellSize int
	Cells    map[GridKey]*Chunk
}

type GridKey struct {
	x, y int
}
