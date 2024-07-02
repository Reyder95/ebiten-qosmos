package ebitenqosmosclasses

// The camera. What defines a camera in a 2d world.
type Camera struct {
	WorldCoordinates Coordinates
}

// Initializes the camera at a specific coordinate
func (c *Camera) InitializeCamera(x, y float64) {
	c.WorldCoordinates.X = x
	c.WorldCoordinates.Y = y
}
