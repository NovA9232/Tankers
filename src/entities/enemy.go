package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
)

type Enemy interface {
	Entity
	Die()
	TakeDamage(int)
	GetIsDead() bool
	GetDimensions() (float32, float32)
}

type BaseEnemy struct {
	Health int
	Damage int
	texture *rl.Texture2D
	isDead bool
}

func (e *BaseEnemy) Draw() {
	rl.DrawTexturePro(*e.texture, rl.NewRectangle(0, 0, e.W, e.H), rl.NewRectangle(e.Pos.X, e.Pos.Y, e.W, e.H), rl.NewVector2(0, 0), 0, rl.Red)
}

func (e *BaseEnemy) Die() {
	e.isDead = true
}

func (e *BaseEnemy) TakeDamage(amount int) {
	e.Health -= amount
	if e.Health < 0 && !e.isDead {
		e.Die()
	}
}

func (e *BaseEnemy) GetIsDead() bool {
	return e.isDead
}
