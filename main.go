// ebitengine Qosmos - A project in GoLang using the ebitengine game framework to build an easy solution for 2D Cameras and Game World Traversing
// Written by Konstantinos Houtas

// TL;DR ebitengine has no easy built in methods to develop game world and movement across a large area. This is why many ebitengine template games are super simple. This project
// tries to build an easy-to-use camera system within this framework. Able to be used by everyone.

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Global coordinates struct. All types of 2d coordinates will use this.
type Coordinates struct {
	x float64
	y float64
}

// Tile. The game world consists of tiles linked together on all 4 sides. This demonstrates the values of a tile
type Tile struct {
	width           float64
	height          float64
	color           color.Color
	realCoordinates Coordinates
	tileImage       *ebiten.Image
	links           TileLink
}

// Initializes a tile with a direct set of coordinates. This is usually done as the first tile. The remaining tiles are built around it.
func (t *Tile) initializeTileWithCoords(_tileHeight, _tileWidth float64, _color color.Color, _realCoordinates Coordinates) {
	t.height = _tileHeight
	t.width = _tileWidth
	t.color = _color
	t.realCoordinates = _realCoordinates
}

// Initializes a tile without a set of coordinates. This is used for every tile that isn't the first because that tile's coordinates is decided based on links and tile size
func (t *Tile) initializeTile(_tileHeight, _tileWidth float64, _color color.Color) {
	t.height = _tileHeight
	t.width = _tileWidth
	t.color = _color
}

// Generates an image for the specific tile based on the tile's color. This will be changed in the future for sprites
func (t *Tile) generateImage() {
	t.tileImage = ebiten.NewImage(int(t.width), int(t.height))
	t.tileImage.Fill(t.color)
}

// Sets a tile offset from another tile in the specified direction. Starting with _tile1. e.g. If I provide Tile1, Tile2, "top", then it will put tile2 on top of tile1.
func (t *Tile) linkTiles(_tile1, _tile2 *Tile, direction string) {
	if direction == "top" {
		_tile1.links.top = _tile2
		_tile2.links.bottom = _tile1
		_tile2.realCoordinates.x = _tile1.realCoordinates.x
		_tile2.realCoordinates.y = _tile1.realCoordinates.y + _tile1.height
	}

	if direction == "bottom" {
		_tile1.links.bottom = _tile2
		_tile2.links.top = _tile1
		_tile2.realCoordinates.x = _tile1.realCoordinates.x
		_tile2.realCoordinates.y = _tile1.realCoordinates.y - _tile1.height
	}

	if direction == "left" {
		_tile1.links.left = _tile2
		_tile2.links.right = _tile1
		_tile2.realCoordinates.x = _tile1.realCoordinates.x + _tile1.width
		_tile2.realCoordinates.y = _tile1.realCoordinates.y
	}

	if direction == "right" {
		_tile1.links.right = _tile2
		_tile2.links.left = _tile1
		_tile2.realCoordinates.x = _tile1.realCoordinates.x - _tile1.width
		_tile2.realCoordinates.y = _tile1.realCoordinates.y
	}
}

// The definition of a link. Tiles are linked together, this is how that is done
type TileLink struct {
	top    *Tile
	bottom *Tile
	left   *Tile
	right  *Tile
}

// The game world entirely. A list of all tiles, and the global width and height of the world
type GameWorld struct {
	width    float64
	height   float64
	tileList []Tile
}

func (g *GameWorld) initializeGameWorld(_worldWidth, _worldHeight float64, _tileList []Tile) {
	g.width = _worldWidth
	g.height = _worldHeight
	g.tileList = _tileList
}

// The camera. What defines a camera in a 2d world.
type Camera struct {
	worldCoordinates Coordinates
}

// Initializes the camera at a specific coordinate
func (c *Camera) initializeCamera(x, y float64) {
	c.worldCoordinates.x = x
	c.worldCoordinates.y = y
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
	mainCamera Camera    // The main camera that the user will "see out of". The future idea will be the ability to swap main cameras so you can view in different ways if you want
	gameWorld  GameWorld // The game world of the current Qosmos project
)

func (g *Game) Update() error {

	// -- Basic camera movement controls
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		mainCamera.worldCoordinates.y -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		mainCamera.worldCoordinates.y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		mainCamera.worldCoordinates.x -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		mainCamera.worldCoordinates.x += 5
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
		tileCoords := CalculateScreenCoordinates(gameWorld.tileList[i])
		op.GeoM.Translate(tileCoords.x, tileCoords.y)
		screen.DrawImage(gameWorld.tileList[i].tileImage, op)
		i++
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (_screenWidth, _screenHeight int) {
	return screenWidth, screenHeight
}

func main() {
	mainCamera.initializeCamera(0, 0) // Initialize the camera at coordinate (0, 0) of the game world

	var tileWorld []Tile // Global slice of all tiles that will be a part of this world
	var currTile Tile    // The current tile we are modifying and adding to the tile world

	// Initialize the "starter" tile at position (0, 0) and generate the image for that tile (should maybe move generateImage into the tile initialization function)
	currTile.initializeTileWithCoords(tileSize, tileSize, color.White, Coordinates{x: 0.0, y: 0.0})
	currTile.generateImage()

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
		currTile.initializeTile(tileSize, tileSize, tileColor)
		currTile.linkTiles(&tileWorld[len(tileWorld)-1], &currTile, "top")
		currTile.generateImage()

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

// calculates the coordinates on screen of the current tile we want to display. Based on the current "camera" position.
func CalculateScreenCoordinates(_tile Tile) Coordinates {
	return Coordinates{x: screenWidth/2 - (_tile.realCoordinates.x + (_tile.width / 2)) - mainCamera.worldCoordinates.x, y: screenHeight/2 - (_tile.realCoordinates.y + (_tile.height / 2)) - mainCamera.worldCoordinates.y}
}
