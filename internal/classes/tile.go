package ebitenqosmosclasses

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Tile. The game world consists of tiles linked together on all 4 sides. This demonstrates the values of a tile
type Tile struct {
	Width           float64
	Height          float64
	Color           color.Color
	RealCoordinates Coordinates
	TileImage       *ebiten.Image
	Links           TileLink
	calculated      bool
}

// The definition of a link. Tiles are linked together, this is how that is done
type TileLink struct {
	top    *Tile
	bottom *Tile
	left   *Tile
	right  *Tile
}

// Initializes a tile with a direct set of coordinates. This is usually done as the first tile. The remaining tiles are built around it.
func (t *Tile) InitializeTileWithCoords(_tileHeight, _tileWidth float64, _color color.Color, _realCoordinates Coordinates) {
	t.Height = _tileHeight
	t.Width = _tileWidth
	t.Color = _color
	t.RealCoordinates = _realCoordinates
	t.calculated = true
}

// Initializes a tile without a set of coordinates. This is used for every tile that isn't the first because that tile's coordinates is decided based on links and tile size
func (t *Tile) InitializeTile(_tileHeight, _tileWidth float64, _color color.Color) {
	t.Height = _tileHeight
	t.Width = _tileWidth
	t.Color = _color
	t.calculated = false
}

// Generates an image for the specific tile based on the tile's color. This will be changed in the future for sprites
func (t *Tile) GenerateImage() {
	t.TileImage = ebiten.NewImage(int(t.Width), int(t.Height))
	t.TileImage.Fill(t.Color)
}

// Sets a tile offset from another tile in the specified direction. Starting with _tile1. e.g. If I provide Tile1, Tile2, "top", then it will put tile2 on top of tile1.
func (t *Tile) LinkTiles(_tile1, _tile2 *Tile, direction string) {

	if !t.calculated {
		if direction == "top" {
			_tile1.Links.top = _tile2
			_tile2.Links.bottom = _tile1
			_tile2.RealCoordinates.X = _tile1.RealCoordinates.X
			_tile2.RealCoordinates.Y = _tile1.RealCoordinates.Y + _tile1.Height

		}

		if direction == "bottom" {
			if _tile1.Links.bottom != _tile2 {
				_tile1.Links.bottom = _tile2
				_tile2.Links.top = _tile1
				_tile2.RealCoordinates.X = _tile1.RealCoordinates.X
				_tile2.RealCoordinates.Y = _tile1.RealCoordinates.Y - _tile1.Height
			}

		}

		if direction == "left" {
			if _tile1.Links.left != _tile2 {
				_tile1.Links.left = _tile2
				_tile2.Links.right = _tile1
				_tile2.RealCoordinates.X = _tile1.RealCoordinates.X + _tile1.Width
				_tile2.RealCoordinates.Y = _tile1.RealCoordinates.Y
			}

		}

		if direction == "right" {
			_tile1.Links.right = _tile2
			_tile2.Links.left = _tile1
			_tile2.RealCoordinates.X = _tile1.RealCoordinates.X - _tile1.Width
			_tile2.RealCoordinates.Y = _tile1.RealCoordinates.Y

		}
	}

	t.calculated = true
}
