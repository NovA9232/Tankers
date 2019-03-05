package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	G *Game
)

type Entity interface { // Interface for body
  Draw()
  Update(float32)
}

type Projectile interface {
	Entity
	getTimeOfCreation() float32
	getTimeLimit() float32
}

type Entities struct {
	players []*Tank
	projec []Projectile			 // Projectiles
	targets []*Target
	other []Entity
}

func (e *Entities) AddPlayer(pos rl.Vector2) {
	e.players = append(e.players, NewTank(len(e.players), pos))
}

func (e *Entities) AddTarget(pos rl.Vector2, w, h int) {
	e.targets = append(e.targets, NewTarget(pos, w, h))
}

func (e *Entities) DrawAllEntites() {
	for i := 0; i < len(e.players); i++ {
		e.players[i].Draw()
	}
	for i := 0; i < len(e.targets); i++ {
		e.targets[i].Draw()
	}
	for i := 0; i < len(e.projec); i++ {
		e.projec[i].Draw()
	}
	for i := 0; i < len(e.other); i++ {
		e.other[i].Draw()
	}
}

func (e *Entities) UpdateAllEntites(dt float32) {
	for i := 0; i < len(e.players); i++ {
		e.players[i].Update(dt)
	}
	for i := 0; i < len(e.targets); i++ {
		e.targets[i].Update(dt)
	}
	for i := 0; i < len(e.projec); i++ {
		e.projec[i].Update(dt)
	}
	for i := 0; i < len(e.other); i++ {
		e.other[i].Update(dt)
	}

	e.checkProjectileTimeouts()
}

func (e *Entities) checkProjectileTimeouts() {
	currTime := rl.GetTime()
	for i := 0; i < len(e.projec); i++ {
		if currTime - e.projec[i].getTimeOfCreation() > e.projec[i].getTimeLimit() {
			copy(e.projec[i:], e.projec[i+1:])
			e.projec[len(e.projec)-1] = nil
			e.projec = e.projec[:len(e.projec)-1]
		}
	}
}
