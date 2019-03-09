package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"

	"mapping"
	"entities"
	"anim"
)

const (
	SCREEN_W = 1280
	SCREEN_H = 720

	crosshairL = 7   // Length of each part of the crosshair
	crosshairW = 2    // Width of each part of the crosshair. Works best with multiples of 2
	crosshairGap = 3 // Gap between each part of crosshair
)

var (
	crWDiv2 int32 = int32(math.Round(crosshairW/2))
	ent *entities.Entities
)

func main() {
	entities.SCREEN_W = SCREEN_W
	entities.SCREEN_H = SCREEN_H

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Tankers")
	defer rl.CloseWindow()
	rl.SetTargetFPS(144)
	rl.HideCursor()

	G := &entities.Game {
		WorldMap: mapping.TestMap(SCREEN_W, SCREEN_H),
		Ent: new(entities.Entities),
		Anim: []*anim.Animation{},
	}
	entities.G = G
	G.Ent.AddPlayer(rl.NewVector2(400, 500))

	for !rl.WindowShouldClose() {
		G.Update(rl.GetFrameTime())

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		G.Draw()
		drawCrosshair()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
