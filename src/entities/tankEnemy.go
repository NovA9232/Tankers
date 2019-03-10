package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	TANK_ENEMY_HEALTH = 100
)

type TankEnemy struct {
	Tank
	BaseEnemy
}

func NewTankEnemy(IDNum int, pos rl.Vector2) *TankEnemy {
	t := &TankEnemy {
		BaseEnemy: BaseEnemy {
			W: TANK_W,
			H: TANK_H,
			Health: TANK_ENEMY_HEALTH,
			Damage: TANK_SHELL_DAMAGE,
			isDead: false,
		},
		Tank: *NewTank(IDNum, pos, false),
	}
	t.texture = &tankTex
	t.Tank.BaseBody.Id.Type = "enemy"

	return t
}

func (t *TankEnemy) Draw() {
	t.BaseEnemy.Draw()
}

func (t *TankEnemy) Update(dt float32) {
	t.applyResistance(dt)

	vel := tools.GetXYComponent(t.VelMag, t.Angle)
  t.Pos.X += (vel.X*dt)
  t.Pos.Y -= (vel.Y*dt)

	t.Cannon.update(dt)
}
