package ebitenqosmosutils

import (
	"example.com/ebiten-qosmos-classes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// calculates the coordinates on screen of the current tile we want to display. Based on the current "camera" position.
func CalculateScreenCoordinates(mainCamera ebitenqosmosclasses.Camera, screenWidth, screenHeight int, _tile ebitenqosmosclasses.Tile) ebitenqosmosclasses.Coordinates {
	return ebitenqosmosclasses.Coordinates{X: float64(screenWidth)/2 - (_tile.RealCoordinates.X + (_tile.Width / 2)) - mainCamera.WorldCoordinates.X, Y: float64(screenHeight)/2 - (_tile.RealCoordinates.Y + (_tile.Height / 2)) - mainCamera.WorldCoordinates.Y}
}

func NewGridKey(x, y int) ebitenqosmosclasses.GridKey {
	return ebitenqosmosclasses.GridKey{
		X: x,
		Y: y,
	}
}

func GetGridKey(x, y, cellSize int) ebitenqosmosclasses.GridKey {
	return ebitenqosmosclasses.GridKey{
		X: -(x / cellSize),
		Y: -(y / cellSize),
	}
}

func DrawChunksNearPlayer(screenWidth, screenHeight, cellSize int, screen *ebiten.Image, camera ebitenqosmosclasses.Camera, viewDistance int, chunkGrid map[ebitenqosmosclasses.GridKey]*ebitenqosmosclasses.Chunk) {
	for dx := int(camera.WorldCoordinates.X) - viewDistance; dx < int(camera.WorldCoordinates.X)+viewDistance; {
		for dy := int(camera.WorldCoordinates.Y) - viewDistance; dy < int(camera.WorldCoordinates.Y)+viewDistance; {
			chunkGridKey := GetGridKey(dx, dy, cellSize)
			chunk, found := chunkGrid[chunkGridKey]

			if found {
				fmt.Println("TEST!")
				for i := 0; i < len(chunk.TileList); i++ {
					op := &ebiten.DrawImageOptions{}
					tileCoords := CalculateScreenCoordinates(camera, screenWidth, screenHeight, chunk.TileList[i])
					op.GeoM.Translate(tileCoords.X, tileCoords.Y)
					screen.DrawImage(chunk.TileList[i].TileImage, op)
				}
			}

			dy += cellSize
		}

		dx += cellSize
	}
}
