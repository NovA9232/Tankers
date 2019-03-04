package mapping

import (
	"github.com/gen2brain/raylib-go/raylib"
	"mapping/tiles"
)

//type MapTile interface {
//	Draw()
//}

type Map struct {
	Tiles []*tiles.Tile
	Width int   // In number of tiles
	Height int
	TileW int
	TileH int
}

func (m *Map) Draw() {
	for i := 0; i < len(m.Tiles); i++ {
		m.Tiles[i].Draw()
	}
}

func TestMap(w, h int) *Map {
	m := new(Map)
	m.Width = w
	m.Height = h
	m.TileW = 16
	m.TileH = 16

	for i := 0; i < m.Width; i += m.TileW {
		for j := 0; j< m.Height; j += m.TileH {
			m.Tiles = append(m.Tiles, tiles.NewCrackedStoneTile(rl.NewVector2(float32(i), float32(j))))
		}
	}

	return m
}
