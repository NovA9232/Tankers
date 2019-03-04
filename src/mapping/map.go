package mapping

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
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

func (m *Map) GetResistanceAt(x, y float32) float32 {
	if x < 0 || y < 0 || x > float32(m.Width*m.TileW) || y > float32(m.Height*m.TileH) {
		return 1
	}

	xTileNum := int(math.Floor(float64(x/float32(m.TileW))))
	yTileNum := int(math.Floor(float64(y/float32(m.TileW))))
	tileNum := int((m.Width * yTileNum) + xTileNum)
	if tileNum > len(m.Tiles) {
		return 1
	}

	return m.Tiles[tileNum].Resistance
}

func (m *Map) Draw() {
	for i := 0; i < len(m.Tiles); i++ {
		m.Tiles[i].Draw()
	}
}

func TestMap(w, h int) *Map {
	m := new(Map)
	m.Width = (w/32)+1
	m.Height = (h/32)+1
	m.TileW = 32
	m.TileH = 32

	for i := 0; i < m.Width; i++ {
		for j := 0; j< m.Height; j++ {
			m.Tiles = append(m.Tiles, tiles.NewConcreteTile(rl.NewVector2(float32(i*m.TileW), float32(j*m.TileH))))
		}
	}

	return m
}
