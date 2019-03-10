package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (e *Entities) AddPlayer(pos rl.Vector2) {
	e.players = append(e.players, NewTank(len(e.players), pos, true))
}

func (e *Entities) AddTankEnemy(pos rl.Vector2) {
	e.enemies = append(e.enemies, Enemy( NewTankEnemy(len(e.enemies), pos) ))
}
