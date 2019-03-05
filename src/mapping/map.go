package mapping

import (
	"github.com/gen2brain/raylib-go/raylib"
	"mapping/tiles"
)

//type MapTile interface {
//	Draw()
//}

type Map struct {
	Tiles []tiles.Tile    // First is x, second is y. Each array is a vertical strip of tiles
	Width int   // In number of tiles
	Height int
	TileW int
	TileH int
}

func (m *Map) GetResistanceAt(x, y float32) float32 {
	if x < 0 || y < 0 || x > float32(m.Width*m.TileW) || y > float32(m.Height*m.TileH) {
		return 1
	}

	for i := 0; i < len(m.Tiles); i++ {
		p := m.Tiles[i].GetPos()
		W, H := m.Tiles[i].GetDimensions()
		if p.X < x && p.Y < y && p.X + W > x && p.Y + H > y {  // If in bounds
			return m.Tiles[i].GetResistance()
		}
	}

	return 1
}

func (m *Map) Draw() {
	for i := 0; i < len(m.Tiles); i++ {
		m.Tiles[i].Draw()
	}
}

func (m *Map) Update(dt float32) {
	for i := 0; i < len(m.Tiles); i++ {
		m.Tiles[i].Update(dt)
	}
}

func TestMap(w, h int) *Map {
	m := new(Map)
	m.Width = (w/32)+1
	m.Height = (h/32)+1
	m.TileW = 32
	m.TileH = 32

	m.Tiles = append(m.Tiles, tiles.NewWaterTile(rl.NewVector2(0, 0), float32(m.TileW*5), float32(m.TileH*23)))
	m.Tiles = append(m.Tiles, tiles.NewConcreteTile(rl.NewVector2(float32(m.TileW*5), 0), 1120, 736))

	return m
}
