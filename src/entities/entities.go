package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	PiOv2 float32 = rl.Pi/2
	TwoPi float32 = rl.Pi * 2
)

var (
	SCREEN_W float32
	SCREEN_H float32

	Ent *Entities
)

type Entity interface { // Interface for body
  Draw()
  Update(float32)
}

type Entities struct {
	players []*Tank
	projec []Entity			 // Projectiles
}

func (e *Entities) AddPlayer(pos rl.Vector2) {
	e.players = append(e.players, NewTank(len(e.players), pos))
}

func (e *Entities) DrawAllEntites() {
	for i := 0; i < len(e.projec); i++ {
		e.projec[i].Draw()
	}
	for i := 0; i < len(e.players); i++ {
		e.players[i].Draw()
	}
}

func (e *Entities) UpdateAllEntites(dt float32) {
	for i := 0; i < len(e.projec); i++ {
		e.projec[i].Update(dt)
	}
	for i := 0; i < len(e.players); i++ {
		e.players[i].Update(dt)
	}
}
