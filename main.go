package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"

	"bodies"
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
	bodies.SCREEN_W = SCREEN_W
	bodies.SCREEN_H = SCREEN_H

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Tankers")
	defer rl.CloseWindow()
	rl.HideCursor()
	rl.SetTargetFPS(144)

	backgroundClr := rl.NewColor(58, 48, 43, 255)

	var players []*bodies.Tank

	tank := bodies.NewTank(len(players), rl.NewVector2(200, 200))
	players = append(players, tank)

	var dt float32

	for !rl.WindowShouldClose() {
		dt = rl.GetFrameTime()
		for i := 0; i < len(players); i++ {
			players[i].Update(dt)
		}

		rl.BeginDrawing()
		rl.ClearBackground(backgroundClr)

		for i := 0; i < len(players); i++ {
			players[i].Draw()
		}
		drawCrosshair()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
