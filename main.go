// ebitengine Qosmos - A project in GoLang using the ebitengine game framework to build an easy solution for 2D Cameras and Game World Traversing
// Written by Konstantinos Houtas

// TL;DR ebitengine has no easy built in methods to develop game world and movement across a large area. This is why many ebitengine template games are super simple. This project
// tries to build an easy-to-use camera system within this framework. Able to be used by everyone.

package main

import (
	"fmt"
	"image/color"
	"log"

	"example.com/ebiten-qosmos-classes"
	"example.com/ebiten-qosmos-utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// The game world entirely. A list of all tiles, and the global width and height of the world
type GameWorld struct {
	width    float64
	height   float64
	tileList []ebitenqosmosclasses.Tile
}

func (g *GameWorld) initializeGameWorld(_worldWidth, _worldHeight float64, _tileList []ebitenqosmosclasses.Tile) {
	g.width = _worldWidth
	g.height = _worldHeight
	g.tileList = _tileList
}

type Game struct{}

// Basic consts used across Qosmos. The screen width and height, and the tile size. May turn these into normal vars for customizable use.
const (
	screenWidth  = 1280
	screenHeight = 720
	tileSize     = 6
)

// Variables that are chaneable and assignable
var (
	mainCamera ebitenqosmosclasses.Camera // The main camera that the user will "see out of". The future idea will be the ability to swap main cameras so you can view in different ways if you want
	gameWorld  GameWorld                  // The game world of the current Qosmos project
)

func (g *Game) Update() error {

	// -- Basic camera movement controls
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		mainCamera.WorldCoordinates.Y -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		mainCamera.WorldCoordinates.Y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		mainCamera.WorldCoordinates.X -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		mainCamera.WorldCoordinates.X += 5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Loop through every tile and draw them each to the screen based on the current camera position.
	// This needs to be completely different because if we have thousands of tiles, we want to only draw specific tiles to the screen, not every single tile ever
	i := 0

	for {
		if i > len(gameWorld.tileList)-1 {
			break
		}

		op := &ebiten.DrawImageOptions{}
		tileCoords := ebitenqosmosutils.CalculateScreenCoordinates(mainCamera, screenWidth, screenHeight, gameWorld.tileList[i])
		op.GeoM.Translate(tileCoords.X, tileCoords.Y)
		screen.DrawImage(gameWorld.tileList[i].TileImage, op)
		i++
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (_screenWidth, _screenHeight int) {
	return screenWidth, screenHeight
}

func main() {
	mainCamera.InitializeCamera(-400, -350) // Initialize the camera at coordinate (0, 0) of the game world

	// var tileWorld [][]ebitenqosmosclasses.Tile
	// rowColLimit := 100

	// for j := 0; j < rowColLimit; j++ {
	// 	for i := 0; i < rowColLimit; i++ {

	// 	}
	// }

	// gameWorld.tileList = gameWorld.tileList.append
	var tileWorld [][]ebitenqosmosclasses.Tile
	var tileColor color.Color
	rowColLimit := 100

	for j := 0; j < rowColLimit; j++ {
		var tileRow []ebitenqosmosclasses.Tile
		for i := 0; i < rowColLimit; i++ {
			singleTile := ebitenqosmosclasses.Tile{}

			if (i+j)%2 == 0 {
				tileColor = color.White
			} else {
				tileColor = color.RGBA{R: 100, G: 100, B: 100, A: 255}
			}

			singleTile.InitializeTileWithCoords(tileSize, tileSize, tileColor, ebitenqosmosclasses.Coordinates{X: float64(tileSize * i), Y: float64(tileSize * j)})

			singleTile.GenerateImage()

			fmt.Println("Tile Coords: X: ", float64(tileSize*i), " Y: ", float64(tileSize*j))

			tileRow = append(tileRow, singleTile)

			//gameWorld.tileList = append(gameWorld.tileList, singleTile)
		}

		tileWorld = append(tileWorld, tileRow)
	}

	var worldChunks []ebitenqosmosclasses.Chunk
	chunkSize := 16
	chunkCount := 0

	for startJ := 0; startJ < len(tileWorld); startJ += chunkSize {
		for startI := 0; startI < len(tileWorld[startJ]); startI += chunkSize {
			chunkCount++
			newChunk := ebitenqosmosclasses.Chunk{}

			for j := startJ; j < startJ+chunkSize; j++ {
				if j >= len(tileWorld) {
					break
				}

				for i := startI; i < startI+chunkSize; i++ {
					if i >= len(tileWorld[j]) {
						break
					}
					newChunk.TileList = append(newChunk.TileList, tileWorld[j][i])
					tileWorld[j][i].Color = color.RGBA{R: uint8(30 * chunkCount), G: uint8(30 * chunkCount), B: uint8(30 * chunkCount), A: 255}
					tileWorld[j][i].GenerateImage()
					gameWorld.tileList = append(gameWorld.tileList, tileWorld[j][i])
				}
			}

			worldChunks = append(worldChunks, newChunk)
		}
	}

	// Example output
	for _, chunk := range worldChunks {
		fmt.Printf("Chunk with %d tiles\n", len(chunk.TileList))
	}

	// fmt.Println(len(gameWorld.tileList))

	// //Example output
	for _, chunk := range worldChunks {
		fmt.Printf("Chunk with %d tiles\n", len(chunk.TileList))
	}

	// ebiten setting basic window options
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
