package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"

	"entities"
	"anim"
)

const (
	SCREEN_W = 1200
	SCREEN_H = 1000

	crosshairL = 7   // Length of each part of the crosshair
	crosshairW = 2    // Width of each part of the crosshair. Works best with multiples of 2
	crosshairGap = 3 // Gap between each part of crosshair
)

var (
	crWDiv2 int32 = int32(math.Round(crosshairW/2))
	ent *entities.Entities
)

func drawCrosshair() {
	mouseX, mouseY := rl.GetMouseX(), rl.GetMouseY()
	posYMid := mouseY - crWDiv2
	posXMid := mouseX - crWDiv2

	rl.DrawRectangle(mouseX - crosshairGap - crosshairL, posYMid, crosshairL, crosshairW, rl.Red)   // Left part
  rl.DrawRectangle(mouseX + crosshairGap, posYMid, crosshairL, crosshairW, rl.Red)								// Right part

	rl.DrawRectangle(posXMid, mouseY - crosshairGap - crosshairL, crosshairW, crosshairL, rl.Red)   // Top part
	rl.DrawRectangle(posXMid, mouseY + crosshairGap, crosshairW, crosshairL, rl.Red)								// Bottom part
}

func main() {
	entities.SCREEN_W = SCREEN_W
	entities.SCREEN_H = SCREEN_H

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Tankers")
	defer rl.CloseWindow()
	rl.HideCursor()
	rl.SetTargetFPS(144)

	G := &entities.Game {
		Ent: new(entities.Entities),
		Anim: []*anim.Animation{},
	}
	entities.G = G
	G.Ent.AddPlayer(rl.NewVector2(400, 500))

	for !rl.WindowShouldClose() {
		G.Update(rl.GetFrameTime())

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		G.Draw()
		drawCrosshair()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
