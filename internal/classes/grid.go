package ebitenqosmosclasses

type Grid struct {
	CellSize int
	Cells    map[GridKey]*Chunk
}

type GridKey struct {
	X, Y int
}
