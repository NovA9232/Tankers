package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (e *Entities) AddPlayer(pos rl.Vector2) {
	e.players = append(e.players, NewTank(len(e.players), pos))
}

func (e *Entities) AddBarrelExplosive(pos rl.Vector2) {
	e.explosives = append(e.explosives, Explosive( NewBarrelExplosive(pos) ))
}

func (e *Entities) AddProximityExplosive(pos rl.Vector2) {
	e.explosives = append(e.explosives, Explosive( NewProximityExplosive(pos) ))
}
