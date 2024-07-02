// ebitengine Qosmos - A project in GoLang using the ebitengine game framework to build an easy solution for 2D Cameras and Game World Traversing
// Written by Konstantinos Houtas

// TL;DR ebitengine has no easy built in methods to develop game world and movement across a large area. This is why many ebitengine template games are super simple. This project
// tries to build an easy-to-use camera system within this framework. Able to be used by everyone.

package main

import (
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
	tileSize     = 30
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
	mainCamera.InitializeCamera(0, 0) // Initialize the camera at coordinate (0, 0) of the game world

	var tileWorld []ebitenqosmosclasses.Tile // Global slice of all tiles that will be a part of this world
	var currTile ebitenqosmosclasses.Tile    // The current tile we are modifying and adding to the tile world

	// Initialize the "starter" tile at position (0, 0) and generate the image for that tile (should maybe move generateImage into the tile initialization function)
	currTile.InitializeTileWithCoords(tileSize, tileSize, color.White, ebitenqosmosclasses.Coordinates{X: 0.0, Y: 0.0})
	currTile.GenerateImage()

	tileWorld = append(tileWorld, currTile) // Add the tile to the game world

	i := 0
	for {
		// Set the tile color based on if "i" is even or not. This is purely just for viewing different tiles in the game world.
		var tileColor color.Color

		if i%2 == 0 {
			tileColor = color.White
		} else {
			tileColor = color.RGBA{R: 10, G: 30, B: 100}
		}

		// Initialize the next tile, link the previous tile to the next tile, and generate the image for that tile. This will need to be redone depending on what is sent to this module. Such as a 2d slice of images or elements.
		currTile.InitializeTile(tileSize, tileSize, tileColor)
		currTile.LinkTiles(&tileWorld[len(tileWorld)-1], &currTile, "top")
		currTile.GenerateImage()

		// Add the current tile to the game world
		tileWorld = append(tileWorld, currTile)

		if i > 9 {
			break
		}

		i++
	}

	gameWorld.initializeGameWorld(10000, 10000, tileWorld) // Initialize the game world with a width of 10,000 and a height of 10,000. Send in the slice of tiles.

	// ebiten setting basic window options
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
