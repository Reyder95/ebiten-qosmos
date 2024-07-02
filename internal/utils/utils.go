package ebitenqosmosutils

import (
	"example.com/ebiten-qosmos-classes"
)

// calculates the coordinates on screen of the current tile we want to display. Based on the current "camera" position.
func CalculateScreenCoordinates(mainCamera ebitenqosmosclasses.Camera, screenWidth, screenHeight int, _tile ebitenqosmosclasses.Tile) ebitenqosmosclasses.Coordinates {
	return ebitenqosmosclasses.Coordinates{X: float64(screenWidth)/2 - (_tile.RealCoordinates.X + (_tile.Width / 2)) - mainCamera.WorldCoordinates.X, Y: float64(screenHeight)/2 - (_tile.RealCoordinates.Y + (_tile.Height / 2)) - mainCamera.WorldCoordinates.Y}
}
