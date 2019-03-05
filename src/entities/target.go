package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
	b2 "github.com/neguse/go-box2d-lite/box2dlite"
)

type Target struct {
	b2.Body
}

func NewTarget(pos rl.Vector2, w, h int) *Target {
	t := new(Target)
	t.Body.Set(&b2.Vec2{float64(w), float64(h)}, 100000)
	t.Position.X, t.Position.Y = float64(pos.X), float64(pos.Y)
	t.Friction = 0

	G.PhysicsWorld.AddBody(&t.Body)
	return t
}

func (t *Target) Draw() {
	rl.DrawRectangleRec(rl.NewRectangle(float32(t.Position.X-(t.Width.X/2)), float32(t.Position.Y-(t.Width.Y/2)), float32(t.Width.X), float32(t.Width.Y)), rl.Red)
}

func (t *Target) Update(dt float32) {
	t.Velocity.X, t.Velocity.Y = 0, 0
}
